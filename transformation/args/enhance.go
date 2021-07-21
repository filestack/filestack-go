package args

// EnhanceArgs args for Enhance transformation.
type EnhanceArgs struct {
	preset *string
}

// NewEnhanceArgs constructor.
func NewEnhanceArgs() *EnhanceArgs {
	return &EnhanceArgs{}
}

// SetPreset setter.
func (a *EnhanceArgs) SetPreset(preset string) *EnhanceArgs {
	a.preset = &preset
	return a
}

// ToMap converts this data to a map.
func (a *EnhanceArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.preset != nil {
		args["preset"] = a.preset
	}

	return args
}
