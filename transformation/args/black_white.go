package args

// BlackWhiteArgs args for BlackWhite transformation.
type BlackWhiteArgs struct {
	threshold *int
}

// NewBlackWhiteArgs constructor.
func NewBlackWhiteArgs() *BlackWhiteArgs {
	return &BlackWhiteArgs{}
}

// SetThreshold setter.
func (a *BlackWhiteArgs) SetThreshold(threshold int) *BlackWhiteArgs {
	a.threshold = &threshold
	return a
}

// ToMap converts this data to a map.
func (a *BlackWhiteArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.threshold != nil {
		args["threshold"] = a.threshold
	}

	return args
}
