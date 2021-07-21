package args

// SmartCropArgs args for SmartCrop transformation.
type SmartCropArgs struct {
	mode      *string
	width     *int
	height    *int
	fillColor *string
	coords    *bool
}

// NewSmartCropArgs constructor.
func NewSmartCropArgs() *SmartCropArgs {
	return &SmartCropArgs{}
}

// SetMode setter.
func (a *SmartCropArgs) SetMode(mode string) *SmartCropArgs {
	a.mode = &mode
	return a
}

// SetWidth setter.
func (a *SmartCropArgs) SetWidth(width int) *SmartCropArgs {
	a.width = &width
	return a
}

// SetHeight setter.
func (a *SmartCropArgs) SetHeight(height int) *SmartCropArgs {
	a.height = &height
	return a
}

// SetFillColor setter.
func (a *SmartCropArgs) SetFillColor(color string) *SmartCropArgs {
	a.fillColor = &color
	return a
}

// SetCoords setter.
func (a *SmartCropArgs) SetCoords(value bool) *SmartCropArgs {
	a.coords = &value
	return a
}

// ToMap converts this data to a map.
func (a *SmartCropArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.mode != nil {
		args["mode"] = a.mode
	}

	if a.width != nil {
		args["width"] = a.width
	}

	if a.height != nil {
		args["height"] = a.height
	}

	if a.fillColor != nil {
		args["fill_color"] = a.fillColor
	}

	if a.coords != nil {
		args["coords"] = a.coords
	}

	return args
}
