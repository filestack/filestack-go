package payload

// WorkflowStatusResponse stores response of CheckWorkflowStatus method.
type WorkflowStatusResponse struct {
	CreatedAt string `json:"createdAt"`
	JobID     string `json:"jobid"`
	Results   map[string]struct {
		MimeType string `json:"mimetype"`
		Size     int    `json:"size"`
		URL      string `json:"url"`
	} `json:"results"`
	Sources    []string `json:"sources"`
	Status     string   `json:"status"`
	TTL        int      `json:"ttl"`
	UpdatedAt  string   `json:"updatedAt"`
	WorkflowId string   `json:"workflow"`
}
