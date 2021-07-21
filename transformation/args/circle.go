package args

// CircleArgs args for Circle transformation.
type CircleArgs struct {
	background *string
}

// NewCircleArgs constructor.
func NewCircleArgs() *CircleArgs {
	return &CircleArgs{}
}

// SetBackground setter.
func (a *CircleArgs) SetBackground(background string) *CircleArgs {
	a.background = &background
	return a
}

// ToMap converts this data to a map.
func (a *CircleArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.background != nil {
		args["background"] = a.background
	}

	return args
}
