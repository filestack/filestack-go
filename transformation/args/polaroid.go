package args

// PolaroidArgs args for Polaroid transformation.
type PolaroidArgs struct {
	rotate     *int
	background *string
	color      *string
}

// NewPolaroidArgs constructor.
func NewPolaroidArgs() *PolaroidArgs {
	return &PolaroidArgs{}
}

// SetRotate setter.
func (a *PolaroidArgs) SetRotate(rotate int) *PolaroidArgs {
	a.rotate = &rotate
	return a
}

// SetBackground setter.
func (a *PolaroidArgs) SetBackground(background string) *PolaroidArgs {
	a.background = &background
	return a
}

// SetColor setter.
func (a *PolaroidArgs) SetColor(color string) *PolaroidArgs {
	a.color = &color
	return a
}

// ToMap converts this data to a map.
func (a *PolaroidArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.rotate != nil {
		args["rotate"] = a.rotate
	}

	if a.background != nil {
		args["background"] = a.background
	}

	if a.color != nil {
		args["color"] = a.color
	}

	return args
}
