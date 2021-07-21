package intelligent

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/filestack/filestack-go/internal/api"
	"github.com/filestack/filestack-go/internal/api/payload"
	"github.com/filestack/filestack-go/internal/filelink"
	"github.com/filestack/filestack-go/internal/segment"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
)

func Upload(
	ctx context.Context,
	file io.ReadSeeker,
	apiKey string,
	defaultStorage string,
	params *store.Params,
	security *security.Security,
	requestHandler *api.RequestHandler,
	defaultFileName string,
	defaultMimeType string,
	defaultChunkSize int64,
	minChunkSize int64,
	cdnHost string,
	defaultPartSize int64,
	maxConcurrentUploads int,
	maxCompleteRetries int,
) (*filelink.FileLink, error) {
	if params == nil {
		params = &store.Params{}
	}

	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, fmt.Errorf("cannot check file size: %w", err)
	}

	fileName := defaultFileName
	if params.FileName != "" {
		fileName = params.FileName
	}

	mimeType := defaultMimeType
	if params.MimeType != "" {
		mimeType = params.MimeType
	}

	storage := defaultStorage
	if params.Location != "" {
		storage = params.Location
	}

	startPayload := payload.MultipartStartRequest{
		ApiKey:      apiKey,
		FileName:    fileName,
		MimeType:    mimeType,
		Size:        int(size),
		Intelligent: true,
		Store: payload.Store{
			Location:  storage,
			Path:      params.Path,
			Region:    params.Region,
			Container: params.Container,
			Access:    params.Access,
		},
	}

	if security != nil {
		startPayload.Policy = security.PolicyB64
		startPayload.Signature = security.Signature
	}

	startResponse, err := requestHandler.MultipartStart(ctx, startPayload)
	if err != nil {
		return nil, fmt.Errorf("making multipart request has failed: %w", err)
	}

	uploadPayload := payload.MultipartUploadRequest{
		ApiKey:   startPayload.ApiKey,
		URI:      startResponse.URI,
		Region:   startResponse.Region,
		UploadID: startResponse.UploadID,
		Store: payload.Store{
			Location: storage,
		},
		Intelligent: true,
	}

	err = uploadParts(
		ctx,
		file,
		size,
		defaultPartSize,
		defaultChunkSize,
		minChunkSize,
		uploadPayload,
		startResponse.LocationURL,
		maxConcurrentUploads,
		requestHandler,
	)
	if err != nil {
		return nil, fmt.Errorf("uploading parts has failed: %w", err)
	}

	payloadComplete := payload.MultipartCompleteRequest{
		ApiKey:      apiKey,
		URI:         startResponse.URI,
		Region:      startResponse.Region,
		UploadID:    startResponse.UploadID,
		Size:        startPayload.Size,
		FileName:    startPayload.FileName,
		MimeType:    startPayload.MimeType,
		Store:       startPayload.Store,
		Intelligent: true,
	}
	if security != nil {
		payloadComplete.Policy = security.PolicyB64
		payloadComplete.Signature = security.Signature
	}

	if params.WorkFlows != "" {
		payloadComplete.Store.WorkFlows = params.WorkFlows
	}
	if len(params.UploadTags) > 0 {
		payloadComplete.UploadTags = params.UploadTags
	}

	completeResponse, err := requestHandler.MultipartCompleteWithRetry(
		ctx,
		startResponse.LocationURL,
		payloadComplete,
		maxCompleteRetries,
	)

	if err != nil {
		return nil, fmt.Errorf("sending `complete` response has failed: %w", err)
	}
	fileLink, err := filelink.New(cdnHost, completeResponse.Handle, apiKey, security, requestHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to create a filelink: %w", err)
	}

	return fileLink, nil
}

func uploadParts(
	ctx context.Context,
	file io.ReadSeeker,
	fileSize int64,
	partSize int64,
	chunkSize int64,
	minChunkSize int64,
	payload payload.MultipartUploadRequest,
	locationURL string,
	workersNumber int,
	requestHandler *api.RequestHandler,
) error {
	parts, err := segment.MakeSegments(file, partSize)
	if err != nil {
		return fmt.Errorf("making parts has failed: %w", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var jobs []uploadPartJob
	var mutex sync.Mutex
	for i, part := range parts {
		partNumber := i + 1
		jobs = append(jobs, newUploadPartJob(
			ctx,
			&mutex,
			partNumber,
			part,
			fileSize,
			chunkSize,
			minChunkSize,
			payload,
			locationURL,
			requestHandler,
		))
	}

	jobsCh := prepareJobs(ctx.Done(), jobs)

	workers := make([]<-chan uploadPartResult, workersNumber)
	for i := 0; i < workersNumber; i++ {
		workers[i] = uploadPart(ctx.Done(), jobsCh)
	}

	outCh := mergeUploadPartResults(ctx.Done(), workers...)
	for uploadResult := range outCh {
		if uploadResult.isSuccessful() == false {
			return fmt.Errorf("merging upload parts has failed: %w", uploadResult.getError())
		}
	}

	return nil
}

type uploadPartJob struct {
	ctx            context.Context
	mutex          *sync.Mutex
	part           segment.Segment
	fileSize       int64
	chunkSize      int64
	minChunkSize   int64
	payload        payload.MultipartUploadRequest
	locationURL    string
	requestHandler *api.RequestHandler
}

func newUploadPartJob(
	ctx context.Context,
	mutex *sync.Mutex,
	partNumber int,
	part segment.Segment,
	fileSize int64,
	chunkSize int64,
	minChunkSize int64,
	payload payload.MultipartUploadRequest,
	locationURL string,
	requestHandler *api.RequestHandler,
) uploadPartJob {
	payload.Part = partNumber
	return uploadPartJob{
		ctx:            ctx,
		mutex:          mutex,
		part:           part,
		fileSize:       fileSize,
		chunkSize:      chunkSize,
		minChunkSize:   minChunkSize,
		payload:        payload,
		locationURL:    locationURL,
		requestHandler: requestHandler,
	}
}

func (u *uploadPartJob) handle() uploadPartResult {
	partOffset := u.part.Offset
	chunks := u.part.Split(u.chunkSize)
	for len(chunks) > 0 {
		chunk := chunks[0]

		u.mutex.Lock()
		chunkData, err := chunk.Bytes()
		u.mutex.Unlock()

		if err != nil {
			return newUploadPartResultError(fmt.Errorf("reading data from chunk has failed: %w", err))
		}

		u.payload.Size = chunk.Length
		hash := md5.Sum(chunkData)
		u.payload.MD5 = base64.StdEncoding.EncodeToString(hash[:])
		u.payload.Offset = chunk.Offset - partOffset

		uploadResponse, err := u.requestHandler.MultipartUpload(u.ctx, u.locationURL, u.payload)
		if err != nil {
			return newUploadPartResultError(fmt.Errorf("chunk upload has failed: %w", err))
		}

		_, err = u.requestHandler.PutChunk(u.ctx, uploadResponse.URL, uploadResponse.Headers, chunkData)
		if err != nil {
			u.part.Offset = chunk.Offset
			u.chunkSize = u.chunkSize / 2
			if u.chunkSize < u.minChunkSize {
				return newUploadPartResultError(errors.New("reached minimum chunk size"))
			}
			chunks = u.part.Split(u.chunkSize)
			continue
		}
		chunks = chunks[1:]
	}

	commitPayload := payload.MultipartCommitRequest{
		ApiKey:   u.payload.ApiKey,
		URI:      u.payload.URI,
		Region:   u.payload.Region,
		UploadID: u.payload.UploadID,
		Store:    u.payload.Store,
		Size:     u.fileSize,
		Part:     u.payload.Part,
	}
	err := u.requestHandler.MultipartCommit(u.ctx, u.locationURL, commitPayload)
	if err != nil {
		return newUploadPartResultError(fmt.Errorf("making a part commit has failed: %w", err))
	}

	return newUploadPartResultSuccess()
}

type uploadPartResult struct {
	err error
}

func newUploadPartResultSuccess() uploadPartResult {
	return uploadPartResult{}
}

func newUploadPartResultError(err error) uploadPartResult {
	return uploadPartResult{err: err}
}

func (u uploadPartResult) isSuccessful() bool {
	return u.err == nil
}

func (u uploadPartResult) getError() error {
	return u.err
}

func mergeUploadPartResults(done <-chan struct{}, channels ...<-chan uploadPartResult) <-chan uploadPartResult {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	out := make(chan uploadPartResult)
	multiplex := func(ch <-chan uploadPartResult) {
		defer wg.Done()
		for c := range ch {
			select {
			case <-done:
				return
			case out <- c:
			}
		}
	}

	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func uploadPart(done <-chan struct{}, jobs <-chan uploadPartJob) <-chan uploadPartResult {
	output := make(chan uploadPartResult)
	go func() {
		defer close(output)
		for job := range jobs {
			select {
			case <-done:
				return
			case output <- job.handle():
			}
		}
	}()
	return output
}

func prepareJobs(done <-chan struct{}, jobs []uploadPartJob) <-chan uploadPartJob {
	output := make(chan uploadPartJob)
	go func() {
		defer close(output)
		for _, job := range jobs {
			select {
			case <-done:
				return
			case output <- job:
			}
		}
	}()
	return output
}
