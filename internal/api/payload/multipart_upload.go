package payload

// MultipartUploadRequest stores request data for MultipartUpload method.
type MultipartUploadRequest struct {
	ApiKey      string `json:"apikey"`
	Part        int    `json:"part"`
	Size        int64  `json:"size"`
	MD5         string `json:"md5"`
	URI         string `json:"uri"`
	Region      string `json:"region"`
	UploadID    string `json:"upload_id"`
	Store       Store  `json:"store"`
	Intelligent bool   `json:"fii"`
	Offset      int64  `json:"offset,omitempty"`
}

// MultipartUploadResponse stores response of MultipartUpload method.
type MultipartUploadResponse struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}
