package payload

// MultipartCompleteRequest stores request data for MultipartComplete
// and MultipartCompleteWithRetry methods.
type MultipartCompleteRequest struct {
	ApiKey      string   `json:"apikey"`
	URI         string   `json:"uri"`
	Region      string   `json:"region"`
	UploadID    string   `json:"upload_id"`
	Size        int      `json:"size"`
	FileName    string   `json:"filename"`
	MimeType    string   `json:"mimetype"`
	Store       Store    `json:"store"`
	UploadTags  []string `json:"upload_tags"`
	Policy      string   `json:"policy,omitempty"`
	Signature   string   `json:"signature,omitempty"`
	Parts       []S3Part `json:"parts,omitempty"`
	Intelligent bool     `json:"fii,omitempty"`
}

// MultipartCompleteResponse stores response of MultipartComplete method.
type MultipartCompleteResponse struct {
	Handle string `json:"handle"`
}
