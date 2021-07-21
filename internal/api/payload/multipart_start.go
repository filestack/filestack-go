package payload

// MultipartStartRequest stores request data for MultipartStart method.
type MultipartStartRequest struct {
	ApiKey      string `json:"apikey"`
	FileName    string `json:"filename"`
	MimeType    string `json:"mimetype"`
	Size        int    `json:"size"`
	Store       Store  `json:"store"`
	Policy      string `json:"policy,omitempty"`
	Signature   string `json:"signature,omitempty"`
	Intelligent bool   `json:"fii"`
}

// MultipartStartResponse stores response of MultipartStart method.
type MultipartStartResponse struct {
	URI         string `json:"uri"`
	Region      string `json:"region"`
	UploadID    string `json:"upload_id"`
	LocationURL string `json:"location_url"`
}
