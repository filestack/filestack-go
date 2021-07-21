package payload

// StoreResponse stores response of Store method.
type StoreResponse struct {
	Filename string `json:"filename""`
	Handle   string `json:"handle"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}
