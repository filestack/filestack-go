package args

// ShadowArgs args for Shadow transformation.
type ShadowArgs struct {
	blur       *float32
	opacity    *int
	vector     []int
	color      *string
	background *string
}

// NewShadowArgs constructor.
func NewShadowArgs() *ShadowArgs {
	return &ShadowArgs{}
}

// SetBlur setter.
func (a *ShadowArgs) SetBlur(blur float32) *ShadowArgs {
	a.blur = &blur
	return a
}

// SetOpacity setter.
func (a *ShadowArgs) SetOpacity(opacity int) *ShadowArgs {
	a.opacity = &opacity
	return a
}

// SetVector setter.
func (a *ShadowArgs) SetVector(x int, y int) *ShadowArgs {
	a.vector = []int{x, y}
	return a
}

// SetColor setter.
func (a *ShadowArgs) SetColor(color string) *ShadowArgs {
	a.color = &color
	return a
}

// SetBackground setter.
func (a *ShadowArgs) SetBackground(background string) *ShadowArgs {
	a.background = &background
	return a
}

// ToMap converts this data to a map.
func (a *ShadowArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.blur != nil {
		args["blur"] = a.blur
	}

	if a.opacity != nil {
		args["opacity"] = a.opacity
	}

	if len(a.vector) == 2 {
		args["vector"] = a.vector
	}

	if a.color != nil {
		args["color"] = a.color
	}

	if a.background != nil {
		args["background"] = a.background
	}

	return args
}
