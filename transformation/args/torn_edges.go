package args

// TornEdgesArgs args for TornEdges transformation.
type TornEdgesArgs struct {
	spread     []int
	background *string
}

// NewTornEdgesArgs constructor.
func NewTornEdgesArgs() *TornEdgesArgs {
	return &TornEdgesArgs{}
}

// SetSpread setter.
func (a *TornEdgesArgs) SetSpread(first int, second int) *TornEdgesArgs {
	a.spread = []int{first, second}
	return a
}

// SetBackground setter.
func (a *TornEdgesArgs) SetBackground(background string) *TornEdgesArgs {
	a.background = &background
	return a
}

// ToMap converts this data to a map.
func (a *TornEdgesArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if len(a.spread) == 2 {
		args["spread"] = a.spread
	}

	if a.background != nil {
		args["background"] = a.background
	}

	return args
}
