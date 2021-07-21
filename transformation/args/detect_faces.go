package args

// DetectFacesArgs arguments for DetectFaces transformation.
type DetectFacesArgs struct {
	maxSize *float32
	minSize *float32
	color   *string
	export  bool
}

// NewDetectFacesArgs constructor.
func NewDetectFacesArgs() *DetectFacesArgs {
	return &DetectFacesArgs{}
}

// SetMaxSize setter.
func (a *DetectFacesArgs) SetMaxSize(maxSize float32) *DetectFacesArgs {
	a.maxSize = &maxSize
	return a
}

// SetMinSize setter.
func (a *DetectFacesArgs) SetMinSize(minSize float32) *DetectFacesArgs {
	a.minSize = &minSize
	return a
}

// SetColor setter.
func (a *DetectFacesArgs) SetColor(color string) *DetectFacesArgs {
	a.color = &color
	return a
}

// Export setter.
func (a *DetectFacesArgs) Export() *DetectFacesArgs {
	a.export = true
	return a
}

// ToMap converts this data to a map.
func (a *DetectFacesArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.maxSize != nil {
		args["maxsize"] = a.maxSize
	}

	if a.minSize != nil {
		args["minsize"] = a.minSize
	}

	if a.color != nil {
		args["color"] = a.color
	}

	if a.export {
		args["export"] = true
	}

	return args
}
