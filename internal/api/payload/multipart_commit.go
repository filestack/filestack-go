package payload

// MultipartCommitRequest stores request data for MultipartCommitRequest method.
type MultipartCommitRequest struct {
	ApiKey   string `json:"apikey"`
	URI      string `json:"uri"`
	Region   string `json:"region"`
	UploadID string `json:"upload_id"`
	Store    Store  `json:"store"`
	Part     int    `json:"part"`
	Size     int64  `json:"size"`
}
