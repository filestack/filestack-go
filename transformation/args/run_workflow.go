package args

// RunWorkflowArgs args for run_workflow transformation.
type RunWorkflowArgs struct {
	id *string
}

// NewRunWorkflowArgs constructor.
func NewRunWorkflowArgs() *RunWorkflowArgs {
	return &RunWorkflowArgs{}
}

// SetID setter.
func (a *RunWorkflowArgs) SetID(id string) *RunWorkflowArgs {
	a.id = &id
	return a
}

// ToMap converts this data to a map.
func (a *RunWorkflowArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.id != nil {
		args["id"] = a.id
	}

	return args
}
