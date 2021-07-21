package payload

// WorkflowRunResponse stores response of RunWorkflow method.
type WorkflowRunResponse struct {
	CreatedAt  string   `json:"createdAt"`
	JobID      string   `json:"jobid"`
	Sources    []string `json:"sources"`
	Status     string   `json:"status"`
	UpdatedAt  string   `json:"updatedAt"`
	WorkflowID string   `json:"workflow"`
}
