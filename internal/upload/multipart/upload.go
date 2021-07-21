package multipart

import (
	"context"
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
	CDNHost string,
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

	payloadStart := payload.MultipartStartRequest{
		ApiKey:   apiKey,
		FileName: fileName,
		MimeType: mimeType,
		Size:     int(size),
		Store: payload.Store{
			Location:  storage,
			Path:      params.Path,
			Region:    params.Region,
			Container: params.Container,
			Access:    params.Access,
		},
	}
	if security != nil {
		payloadStart.Policy = security.PolicyB64
		payloadStart.Signature = security.Signature
	}

	chunks, err := makeChunks(file, defaultChunkSize)
	if err != nil {
		return nil, err
	}

	startResponse, err := requestHandler.MultipartStart(ctx, payloadStart)
	if err != nil {
		return nil, fmt.Errorf("making multipart request has failed: %w", err)
	}

	var mutex sync.Mutex
	chunkSender := newSender(
		ctx,
		apiKey,
		storage,
		startResponse.URI,
		startResponse.LocationURL,
		startResponse.Region,
		startResponse.UploadID,
		minChunkSize,
		requestHandler,
		&mutex,
	)

	s3parts, err := uploadChunks(ctx, chunkSender, chunks)
	if err != nil {
		return nil, err
	}

	payloadComplete := payload.MultipartCompleteRequest{
		ApiKey:   apiKey,
		URI:      startResponse.URI,
		Region:   startResponse.Region,
		UploadID: startResponse.UploadID,
		Size:     payloadStart.Size,
		FileName: payloadStart.FileName,
		MimeType: payloadStart.MimeType,
		Parts:    s3parts,
		Store:    payloadStart.Store,
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
	completeResponse, err := requestHandler.MultipartComplete(ctx, startResponse.LocationURL, payloadComplete)
	if err != nil {
		return nil, fmt.Errorf("sending `complete` response has failed: %w", err)
	}
	fileLink, err := filelink.New(CDNHost, completeResponse.Handle, apiKey, security, requestHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to create a filelink: %w", err)
	}

	return fileLink, nil
}

func makeChunks(file io.ReadSeeker, defaultChunkSize int64) (chunks []chunk, err error) {
	parts, err := segment.MakeSegments(file, defaultChunkSize)
	if err != nil {
		err = fmt.Errorf("making chunks has failed: %w", err)
		return
	}

	for i, part := range parts {
		chunks = append(chunks, chunk{filePart: part, number: i + 1})
	}

	return
}

func uploadChunks(
	ctx context.Context,
	sender *sender,
	chunks []chunk,
) (s3parts []payload.S3Part, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	chunksCh := prepareChunks(ctx.Done(), chunks)

	workersNumber := 5
	workers := make([]<-chan chunkUploadResult, workersNumber)
	for i := 0; i < workersNumber; i++ {
		workers[i] = handleChunkUpload(ctx.Done(), sender, chunksCh)
	}

	outCh := merge(ctx.Done(), workers...)
	for uploadResult := range outCh {
		if uploadResult.isSuccessful() == false {
			err = uploadResult.getError()
			return
		}
		s3parts = append(s3parts, uploadResult.getS3Part())
	}

	return
}

func prepareChunks(done <-chan struct{}, chunks []chunk) <-chan chunk {
	output := make(chan chunk)
	go func() {
		defer close(output)
		for _, c := range chunks {
			select {
			case <-done:
				return
			case output <- c:
			}
		}
	}()
	return output
}

func handleChunkUpload(
	done <-chan struct{},
	sender *sender,
	chunks <-chan chunk,
) <-chan chunkUploadResult {
	output := make(chan chunkUploadResult)
	go func() {
		defer close(output)
		for ch := range chunks {
			s3part, err := sender.send(ch)
			uploadResult := newChunkUploadResult(s3part, err)
			select {
			case <-done:
				return
			case output <- uploadResult:
			}
		}
	}()
	return output
}

func merge(done <-chan struct{}, channels ...<-chan chunkUploadResult) <-chan chunkUploadResult {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	out := make(chan chunkUploadResult)
	multiplex := func(ch <-chan chunkUploadResult) {
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
