package payload

// Store is a type shared in payloads of MultipartCommit,
// MultipartComplete, MultipartStart and MultipartUpload methods.
// It defines storage parameters.
type Store struct {
	Path      string `json:"path,omitempty"`
	Location  string `json:"location"`
	Region    string `json:"region,omitempty"`
	Container string `json:"container,omitempty"`
	Access    string `json:"access,omitempty"`
	WorkFlows string `json:"workflows,omitempty"`
}

// S3Part describes a file chunk.
type S3Part struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}
