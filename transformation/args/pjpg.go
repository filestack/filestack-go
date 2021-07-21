package args

// PJPGArgs args for PJPG transformation.
type PJPGArgs struct {
	quality  *int
	metadata *bool
}

// NewPJPGArgs constructor.
func NewPJPGArgs() *PJPGArgs {
	return &PJPGArgs{}
}

// SetQuality setter.
func (a *PJPGArgs) SetQuality(quality int) *PJPGArgs {
	a.quality = &quality
	return a
}

// SetMetadata setter.
func (a *PJPGArgs) SetMetadata(value bool) *PJPGArgs {
	a.metadata = &value
	return a
}

// ToMap converts this data to a map.
func (a *PJPGArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.quality != nil {
		args["quality"] = a.quality
	}

	if a.metadata != nil {
		args["metadata"] = a.metadata
	}

	return args
}
