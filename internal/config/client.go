package config

import (
	"net/http"

	"github.com/filestack/filestack-go/security"
)

// Client related config parameters.
type Client struct {
	Storage              string
	DefaultFileName      string
	DefaultMimeType      string
	DefaultChunkSize     int64
	DefaultPartSize      int64
	MinChunkSize         int64
	UploadHost           string
	APIHost              string
	CDNHost              string
	MaxConcurrentUploads int
	MaxCompleteRetries   int
	HTTPClient           *http.Client
	Security             *security.Security
}

// NewClientConfig constructor.
func NewClientConfig() *Client {
	return &Client{
		Storage:              "s3",
		DefaultFileName:      "unnamed_file",
		DefaultMimeType:      "application/octet-stream",
		DefaultChunkSize:     5 * 1024 * 1024,
		DefaultPartSize:      8 * 1024 * 1024,
		MinChunkSize:         32 * 1024,
		UploadHost:           "https://upload.filestackapi.com",
		APIHost:              "https://www.filestackapi.com/api",
		CDNHost:              "https://cdn.filestackcontent.com",
		MaxConcurrentUploads: 5,
		MaxCompleteRetries:   7,
		HTTPClient:           http.DefaultClient,
	}
}
