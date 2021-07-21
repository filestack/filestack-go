package payload

// MetadataResponse stores response of GetMetadata method.
type MetadataResponse struct {
	Filename  string  `json:"filename"`
	MimeType  string  `json:"mimetype"`
	Size      int64   `json:"size"`
	Uploaded  float64 `json:"uploaded"`
	Writeable bool    `json:"writeable"`
	Container string  `json:"container"`
	Location  string  `json:"location"`
	Key       string  `json:"key"`
	Path      string  `json:"path"`
}
