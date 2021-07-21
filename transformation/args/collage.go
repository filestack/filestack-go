package args

// CollageArgs args for Collage transformation.
type CollageArgs struct {
	files      []string
	margin     *int
	width      *int
	height     *int
	color      *string
	fit        *string
	autoRotate bool
}

// NewCollageArgs constructor.
func NewCollageArgs() *CollageArgs {
	return &CollageArgs{}
}

// SetFile setter.
func (a *CollageArgs) SetFile(file string) *CollageArgs {
	a.files = []string{file}
	return a
}

// SetFiles setter.
func (a *CollageArgs) SetFiles(files []string) *CollageArgs {
	a.files = files
	return a
}

// SetMargin setter.
func (a *CollageArgs) SetMargin(margin int) *CollageArgs {
	a.margin = &margin
	return a
}

// SetWidth setter.
func (a *CollageArgs) SetWidth(width int) *CollageArgs {
	a.width = &width
	return a
}

// SetHeight setter.
func (a *CollageArgs) SetHeight(height int) *CollageArgs {
	a.height = &height
	return a
}

// SetColor setter.
func (a *CollageArgs) SetColor(color string) *CollageArgs {
	a.color = &color
	return a
}

// SetFit setter.
func (a *CollageArgs) SetFit(fit string) *CollageArgs {
	a.fit = &fit
	return a
}

// SetAutoRotate setter.
func (a *CollageArgs) SetAutoRotate() *CollageArgs {
	a.autoRotate = true
	return a
}

// ToMap converts this data to a map.
func (a *CollageArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if len(a.files) > 0 {
		args["files"] = a.files
	}

	if a.margin != nil {
		args["margin"] = a.margin
	}

	if a.width != nil {
		args["width"] = a.width
	}

	if a.height != nil {
		args["height"] = a.height
	}

	if a.color != nil {
		args["color"] = a.color
	}

	if a.fit != nil {
		args["fit"] = a.fit
	}

	if a.autoRotate {
		args["autorotate"] = true
	}

	return args
}
