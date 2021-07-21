package args

// RoundedCornersArgs args for RoundCorners transformation.
type RoundedCornersArgs struct {
	radiusMax  bool
	radius     *int
	blur       *float32
	background *string
}

// NewRoundedCornersArgs constructor.
func NewRoundedCornersArgs() *RoundedCornersArgs {
	return &RoundedCornersArgs{}
}

// SetRadiusMax setter.
func (a *RoundedCornersArgs) SetRadiusMax() *RoundedCornersArgs {
	a.radiusMax = true
	return a
}

// SetRadius setter.
func (a *RoundedCornersArgs) SetRadius(radius int) *RoundedCornersArgs {
	a.radius = &radius
	return a
}

// SetBlur setter.
func (a *RoundedCornersArgs) SetBlur(blur float32) *RoundedCornersArgs {
	a.blur = &blur
	return a
}

// SetBackground setter.
func (a *RoundedCornersArgs) SetBackground(background string) *RoundedCornersArgs {
	a.background = &background
	return a
}

// ToMap converts this data to a map.
func (a *RoundedCornersArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.radiusMax {
		args["radius"] = "max"
	} else if a.radius != nil {
		args["radius"] = a.radius
	}

	if a.blur != nil {
		args["blur"] = a.blur
	}

	if a.background != nil {
		args["background"] = a.background
	}

	return args
}
