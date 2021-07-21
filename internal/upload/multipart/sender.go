package multipart

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"sync"

	"github.com/filestack/filestack-go/internal/api"
	"github.com/filestack/filestack-go/internal/api/payload"
	"github.com/filestack/filestack-go/internal/segment"
)

// sender is an internal service responsible for file-chunk delivery.
type sender struct {
	ctx            context.Context
	apiKey         string
	storage        string
	uri            string
	locationURL    string
	region         string
	uploadID       string
	chunkSize      int64
	requestHandler *api.RequestHandler
	mutex          *sync.Mutex
}

// newSender constructor.
func newSender(
	ctx context.Context,
	apiKey string,
	storage string,
	uri string,
	locationURL string,
	region string,
	uploadID string,
	chunkSize int64,
	requestHandler *api.RequestHandler,
	mutex *sync.Mutex,
) *sender {
	return &sender{
		ctx:            ctx,
		apiKey:         apiKey,
		storage:        storage,
		uri:            uri,
		locationURL:    locationURL,
		region:         region,
		uploadID:       uploadID,
		chunkSize:      chunkSize,
		requestHandler: requestHandler,
		mutex:          mutex,
	}
}

// Send handles file-chunk delivery.
func (c *sender) send(chunk chunk) (s3part payload.S3Part, err error) {
	c.mutex.Lock()
	chunkData, err := chunk.filePart.Bytes()
	c.mutex.Unlock()
	if err != nil {
		return
	}
	hash := md5.Sum(chunkData)
	hashEncoded := base64.StdEncoding.EncodeToString(hash[:])

	payloadUpload := payload.MultipartUploadRequest{
		ApiKey:   c.apiKey,
		Part:     chunk.number,
		Size:     int64(len(chunkData)),
		MD5:      hashEncoded,
		URI:      c.uri,
		Region:   c.region,
		UploadID: c.uploadID,
		Store: payload.Store{
			Location: c.storage,
		},
	}

	response, err := c.requestHandler.MultipartUpload(c.ctx, c.locationURL, payloadUpload)
	if err != nil {
		return
	}

	etag, err := c.requestHandler.PutChunk(c.ctx, response.URL, response.Headers, chunkData)
	if err != nil {
		return
	}

	s3part = payload.S3Part{
		chunk.number,
		etag,
	}

	return
}

type chunkUploadResult struct {
	s3part payload.S3Part
	err    error
}

func newChunkUploadResult(s3part payload.S3Part, err error) chunkUploadResult {
	return chunkUploadResult{
		s3part: s3part,
		err:    err,
	}
}

func (c chunkUploadResult) isSuccessful() bool {
	return c.err == nil
}

func (c chunkUploadResult) getError() error {
	return c.err
}

func (c chunkUploadResult) getS3Part() payload.S3Part {
	return c.s3part
}

type chunk struct {
	filePart segment.Segment
	number   int
}
