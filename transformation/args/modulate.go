package args

// ModulateArgs args for Modulate transformation.
type ModulateArgs struct {
	brightness *int
	saturation *int
	hue        *int
}

// NewModulateArgs constructor.
func NewModulateArgs() *ModulateArgs {
	return &ModulateArgs{}
}

// SetBrightness setter.
func (a *ModulateArgs) SetBrightness(brightness int) *ModulateArgs {
	a.brightness = &brightness
	return a
}

// SetSaturation setter.
func (a *ModulateArgs) SetSaturation(saturation int) *ModulateArgs {
	a.saturation = &saturation
	return a
}

// SetHue setter.
func (a *ModulateArgs) SetHue(hue int) *ModulateArgs {
	a.hue = &hue
	return a
}

// ToMap converts this data to a map.
func (a *ModulateArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.brightness != nil {
		args["brightness"] = a.brightness
	}

	if a.saturation != nil {
		args["saturation"] = a.saturation
	}

	if a.hue != nil {
		args["hue"] = a.hue
	}

	return args
}
