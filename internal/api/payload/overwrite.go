package payload

// OverwriteResponse stores resonse of Overwrite method.
type OverwriteResponse struct {
	URL         string `json:"url"`
	MimeType    string `json:"mimetype"`
	IsWriteable bool   `json:"isWriteable"`
	Size        int64  `json:"size"`
	Filename    string `json:"filename"`
}
