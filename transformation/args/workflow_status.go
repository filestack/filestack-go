package args

// WorkflowStatusArgs args for workflow_status transformation.
type WorkflowStatusArgs struct {
	jobID *string
}

// NewWorkflowStatusArgs constructor.
func NewWorkflowStatusArgs() *WorkflowStatusArgs {
	return &WorkflowStatusArgs{}
}

// SetJobID setter.
func (a *WorkflowStatusArgs) SetJobID(jobID string) *WorkflowStatusArgs {
	a.jobID = &jobID
	return a
}

// ToMap converts this data to a map.
func (a *WorkflowStatusArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.jobID != nil {
		args["job_id"] = a.jobID
	}

	return args
}
