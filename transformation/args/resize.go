package args

// ResizeArgs arguments for Resize transformation.
type ResizeArgs struct {
	width  *int
	height *int
	fit    *string
	align  *string
	aligns *[]string
	filter *string
}

// NewResizeArgs constructor.
func NewResizeArgs() *ResizeArgs {
	return &ResizeArgs{}
}

// SetWidth setter.
func (a *ResizeArgs) SetWidth(width int) *ResizeArgs {
	a.width = &width
	return a
}

// SetHeight setter.
func (a *ResizeArgs) SetHeight(height int) *ResizeArgs {
	a.height = &height
	return a
}

// SetFit setter.
func (a *ResizeArgs) SetFit(fit string) *ResizeArgs {
	a.fit = &fit
	return a
}

// SetAlign setter.
func (a *ResizeArgs) SetAlign(align string) *ResizeArgs {
	a.align = &align
	return a
}

// SetAligns setter.
func (a *ResizeArgs) SetAligns(vertical, horizontal string) *ResizeArgs {
	aligns := []string{vertical, horizontal}
	a.aligns = &aligns
	return a
}

// SetFilter setter.
func (a *ResizeArgs) SetFilter(filter string) *ResizeArgs {
	a.filter = &filter
	return a
}

// ToMap converts this data to a map.
func (a *ResizeArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.width != nil {
		args["width"] = a.width
	}

	if a.height != nil {
		args["height"] = a.height
	}

	if a.fit != nil {
		args["fit"] = a.fit
	}

	if a.align != nil {
		args["align"] = a.align
	}

	if a.aligns != nil {
		args["align"] = a.aligns
	}

	if a.filter != nil {
		args["filter"] = a.filter
	}

	return args
}
