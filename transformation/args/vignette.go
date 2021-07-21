package args

// VignetteArgs args for Vignette transformation.
type VignetteArgs struct {
	amount     *int
	blurMode   *string
	background *string
}

// NewVignetteArgs constructor.
func NewVignetteArgs() *VignetteArgs {
	return &VignetteArgs{}
}

// SetAmount setter.
func (a *VignetteArgs) SetAmount(amount int) *VignetteArgs {
	a.amount = &amount
	return a
}

// SetBlurMode setter.
func (a *VignetteArgs) SetBlurMode(blurMode string) *VignetteArgs {
	a.blurMode = &blurMode
	return a
}

// SetBackground setter.
func (a *VignetteArgs) SetBackground(background string) *VignetteArgs {
	a.background = &background
	return a
}

// ToMap converts this data to a map.
func (a *VignetteArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.amount != nil {
		args["amount"] = a.amount
	}

	if a.blurMode != nil {
		args["blurmode"] = a.blurMode
	}

	if a.background != nil {
		args["background"] = a.background
	}

	return args
}
