package args

// UpscaleArgs args for Upscale transformation.
type UpscaleArgs struct {
	noise   *string
	upscale *bool
	style   *string
}

// NewUpscaleArgs constructor.
func NewUpscaleArgs() *UpscaleArgs {
	return &UpscaleArgs{}
}

// SetNoise setter.
func (a *UpscaleArgs) SetNoise(noise string) *UpscaleArgs {
	a.noise = &noise
	return a
}

// SetUpscale setter.
func (a *UpscaleArgs) SetUpscale(upscale bool) *UpscaleArgs {
	a.upscale = &upscale
	return a
}

// SetStyle setter.
func (a *UpscaleArgs) SetStyle(style string) *UpscaleArgs {
	a.style = &style
	return a
}

// ToMap converts this data to a map.
func (a *UpscaleArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.noise != nil {
		args["noise"] = a.noise
	}

	if a.upscale != nil {
		args["upscale"] = a.upscale
	}

	if a.style != nil {
		args["style"] = a.style
	}

	return args
}
