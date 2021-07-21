package args

// DocDetection args for DocDetection transformation.
type DocDetection struct {
	coords     *bool
	preprocess *bool
}

// NewDocDetection constructor.
func NewDocDetection() *DocDetection {
	return &DocDetection{}
}

// SetCoords setter.
func (a *DocDetection) SetCoords(value bool) *DocDetection {
	a.coords = &value
	return a
}

// SetPreprocess setter.
func (a *DocDetection) SetPreprocess(value bool) *DocDetection {
	a.preprocess = &value
	return a
}

// ToMap converts this data to a map.
func (a *DocDetection) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.coords != nil {
		args["coords"] = a.coords
	}

	if a.preprocess != nil {
		args["preprocess"] = a.preprocess
	}

	return args
}
