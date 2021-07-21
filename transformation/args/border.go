package args

// BorderArgs args for Border transformation.
type BorderArgs struct {
	color      *string
	background *string
	width      *int
}

// NewBorderArgs constructor.
func NewBorderArgs() *BorderArgs {
	return &BorderArgs{}
}

// SetColor setter.
func (a *BorderArgs) SetColor(color string) *BorderArgs {
	a.color = &color
	return a
}

// SetBackground setter.
func (a *BorderArgs) SetBackground(background string) *BorderArgs {
	a.background = &background
	return a
}

// SetWidth setter.
func (a *BorderArgs) SetWidth(width int) *BorderArgs {
	a.width = &width
	return a
}

// ToMap converts this data to a map.
func (a *BorderArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.color != nil {
		args["color"] = a.color
	}

	if a.background != nil {
		args["background"] = a.background
	}

	if a.width != nil {
		args["width"] = a.width
	}

	return args
}
