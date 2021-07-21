package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/filestack/filestack-go/internal/api/payload"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
)

const maxWaitLimit = time.Second * time.Duration(60)

// RequestHandler is an internal service
// responsible for handling api requests.
type RequestHandler struct {
	uploadHost string
	apiHost    string
	httpClient *http.Client
}

// NewRequestHandler constructor.
func NewRequestHandler(
	uploadHost string,
	apiHost string,
	httpClient *http.Client,
) *RequestHandler {
	return &RequestHandler{
		uploadHost: uploadHost,
		apiHost:    apiHost,
		httpClient: httpClient,
	}
}

// MultipartStart sends `multipart/start` request.
func (r *RequestHandler) MultipartStart(
	ctx context.Context,
	requestPayload payload.MultipartStartRequest,
) (responsePayload payload.MultipartStartResponse, err error) {
	url := r.uploadHost + "/multipart/start"
	err = r.handleRequest(ctx, http.MethodPost, url, requestPayload, &responsePayload)
	if err != nil {
		return
	}
	if len(responsePayload.URI) == 0 {
		err = errors.New("empty URI was received")
		return
	}

	return
}

// MultipartUpload sends `multipart/upload` request.
func (r *RequestHandler) MultipartUpload(
	ctx context.Context,
	locationURL string,
	requestPayload payload.MultipartUploadRequest,
) (responsePayload payload.MultipartUploadResponse, err error) {
	locationURL = "https://" + locationURL + "/multipart/upload"
	err = r.handleRequest(ctx, http.MethodPost, locationURL, requestPayload, &responsePayload)
	if err != nil {
		return
	}
	if len(responsePayload.URL) == 0 {
		err = errors.New("incorrect response received for multipart upload")
		return
	}

	return
}

// MultipartComplete sends `multipart/complete` request.
func (r *RequestHandler) MultipartComplete(
	ctx context.Context,
	locationURL string,
	requestPayload payload.MultipartCompleteRequest,
) (responsePayload payload.MultipartCompleteResponse, err error) {
	locationURL = "https://" + locationURL + "/multipart/complete"
	err = r.handleRequest(ctx, http.MethodPost, locationURL, requestPayload, &responsePayload)

	return
}

// MultipartCommit sends `multipart/commit` request.
func (r *RequestHandler) MultipartCommit(
	ctx context.Context,
	locationURL string,
	requestPayload payload.MultipartCommitRequest,
) error {
	locationURL = "https://" + locationURL + "/multipart/commit"
	return r.handleRequest(ctx, http.MethodPost, locationURL, requestPayload, nil)
}

// MultipartCompleteWithRetry sends `multipart/complete` request
// and retries on failure.
func (r *RequestHandler) MultipartCompleteWithRetry(
	ctx context.Context,
	locationURL string,
	requestPayload payload.MultipartCompleteRequest,
	retries int,
) (payload.MultipartCompleteResponse, error) {
	var (
		responsePayload payload.MultipartCompleteResponse
		retry           int
		startTime       = time.Now()
	)

	locationURL = "https://" + locationURL + "/multipart/complete"
	for {
		time.Sleep(time.Second * time.Duration(retry))

		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(&requestPayload)
		if err != nil {
			return responsePayload, fmt.Errorf("failed to encode the request: %w", err)
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, locationURL, buffer)
		if err != nil {
			return responsePayload, fmt.Errorf("failed to create a new request: %w", err)
		}
		req.Header.Set("Content-type", "application/json")

		res, err := r.httpClient.Do(req)
		if err != nil {
			if retry <= retries {
				retry++
				continue
			}

			return responsePayload, fmt.Errorf("failed to handle the `complete` request: %w", err)
		}

		if res.StatusCode == http.StatusPartialContent {
			return responsePayload, errors.New("partial content has not been committed")
		}

		if res.StatusCode == http.StatusAccepted {
			if time.Now().Sub(startTime) < maxWaitLimit {
				continue
			}
			return responsePayload, errors.New("max wait limit on `accepted` status has been reached")
		}

		if err = json.NewDecoder(res.Body).Decode(&responsePayload); err != nil {
			if retry <= retries {
				retry++
				continue
			}

			return responsePayload, fmt.Errorf("failed to decode response: %w", err)
		}

		return responsePayload, nil
	}
}

// PutChunk handles file-chunk delivery.
func (r *RequestHandler) PutChunk(
	ctx context.Context,
	url string,
	headers map[string]string,
	data []byte,
) (etag string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return
	}

	etag = resp.Header.Get("Etag")
	return
}

// Process handles `process` endpoint communication.
func (r *RequestHandler) Process(
	ctx context.Context,
	host string,
	apiKey string,
	url string,
	storeParams *store.Params,
	security *security.Security,
) (handle string, err error) {
	requestPayload := payload.NewProcessRequest(apiKey, url, storeParams, security)
	url = host + "/process"
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(&requestPayload)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, buffer)
	req.Header.Set("Content-type", "application/json")
	res, e := r.httpClient.Do(req)
	if e != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New("request processing has failed")
		return
	}

	var responseData struct {
		Handle string `json:"handle"`
	}
	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		return
	}
	handle = responseData.Handle

	return
}

// Store is responsible for storing a file.
func (r *RequestHandler) Store(ctx context.Context, url string) (responsePayload payload.StoreResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)

	return
}

// Overwrite allows to overwrite an existing files under a given handle.
func (r *RequestHandler) Overwrite(
	ctx context.Context,
	handle string,
	file io.Reader,
	base64decode bool,
	security *security.Security,
) (payload.OverwriteResponse, error) {
	var responsePayload payload.OverwriteResponse
	requestURL := r.apiHost + "/file/" + handle

	q := url.Values{}
	q.Add("policy", security.PolicyB64)
	q.Add("signature", security.Signature)
	base64decodeStr := "false"
	if base64decode {
		base64decodeStr = "true"
	}
	q.Add("base64decode", base64decodeStr)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, file)
	if err != nil {
		return responsePayload, fmt.Errorf("creating a request has failed: %w", err)
	}
	req.URL.RawQuery = q.Encode()

	res, err := r.httpClient.Do(req)
	if err != nil {
		return responsePayload, fmt.Errorf("processing request has failed: %w", err)
	}

	err = json.NewDecoder(res.Body).Decode(&responsePayload)
	if err != nil {
		return responsePayload, fmt.Errorf("decoding response has failed: %w", err)
	}

	return responsePayload, nil
}

// Delete allows to remove file by a handle string.
func (r *RequestHandler) Delete(
	ctx context.Context,
	handle string,
	apiKey string,
	security *security.Security,
) error {
	apiURL := r.apiHost + "/file/" + handle

	q := url.Values{}
	q.Add("key", apiKey)
	q.Add("policy", security.PolicyB64)
	q.Add("signature", security.Signature)
	query := q.Encode()

	apiURL += "?" + query

	err := r.handleRequest(ctx, http.MethodDelete, apiURL, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete a handle: %s due to error: %w", handle, err)
	}

	return nil
}

// GetMetadata performs `metadata` task.
func (r *RequestHandler) GetMetadata(
	ctx context.Context,
	url string,
) (responsePayload payload.MetadataResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)
	return
}

// GetTags performs `tags` task.
func (r *RequestHandler) GetTags(
	ctx context.Context,
	url string,
) (responsePayload payload.TagsResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)
	return
}

// GetSfw performs `sfw` task.
func (r *RequestHandler) GetSfw(
	ctx context.Context,
	url string,
) (responsePayload payload.SfwResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)
	return
}

// WorkflowRun starts workflow processing.
func (r *RequestHandler) WorkflowRun(
	ctx context.Context,
	url string,
) (responsePayload payload.WorkflowRunResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)
	return
}

// WorkflowStatus checks a status of a workflow.
func (r *RequestHandler) WorkflowStatus(
	ctx context.Context,
	url string,
) (responsePayload payload.WorkflowStatusResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)
	return
}

// AudioVisualInit starts `AVConvert` transformation.
func (r *RequestHandler) AudioVisualInit(
	ctx context.Context,
	url string,
) (responsePayload payload.AudioVisualInitResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)
	return
}

// AudioVisualStatus checks status of the `AVConvert` transformation.
func (r *RequestHandler) AudioVisualStatus(
	ctx context.Context,
	url string,
) (responsePayload payload.AudioVisualStatusResponse, err error) {
	err = r.handleRequest(ctx, http.MethodGet, url, nil, &responsePayload)
	return
}

// GetTransformation allows to fetch the transformation result as bytes.
func (r *RequestHandler) GetTransformation(
	ctx context.Context,
	url string,
) (result []byte, err error) {
	req, reqErr := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if reqErr != nil {
		err = fmt.Errorf("failed to create a request: %w", reqErr)
		return
	}
	res, doErr := r.httpClient.Do(req)
	if doErr != nil {
		err = fmt.Errorf("failed to make a request: %w", doErr)
		return
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("transformation has failed with http code: %d", res.StatusCode)
		return
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		err = fmt.Errorf("failed to read the response: %w", readErr)
		return
	}
	result = body

	return
}

func (r *RequestHandler) handleRequest(
	ctx context.Context,
	method string,
	url string,
	requestPayload interface{},
	responsePayload interface{},
) error {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(&requestPayload)
	req, err := http.NewRequestWithContext(ctx, method, url, buffer)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}
	req.Header.Set("Content-type", "application/json")
	res, err := r.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("processing request has failed (%s %s): %w", method, url, err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request has failed with status code %v", res.StatusCode)
	}
	if responsePayload == nil {
		return nil
	}

	err = json.NewDecoder(res.Body).Decode(responsePayload)
	if err != nil {
		return fmt.Errorf("response decoding has failed: %w", err)
	}

	return nil
}
