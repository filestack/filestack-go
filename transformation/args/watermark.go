package args

// WatermarkArgs arguments for Watermark transformation.
type WatermarkArgs struct {
	file      string
	size      *int
	position  *string
	positions *[]string
}

// NewWatermarkArgs constructor.
func NewWatermarkArgs(file string) *WatermarkArgs {
	return &WatermarkArgs{
		file: file,
	}
}

// SetSize setter.
func (a *WatermarkArgs) SetSize(size int) *WatermarkArgs {
	a.size = &size
	return a
}

// SetPosition setter.
func (a *WatermarkArgs) SetPosition(position string) *WatermarkArgs {
	a.position = &position
	return a
}

// SetPositions setter.
func (a *WatermarkArgs) SetPositions(vertical string, horizontal string) *WatermarkArgs {
	a.positions = &[]string{vertical, horizontal}
	return a
}

// ToMap converts this data to a map.
func (a *WatermarkArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{
		"file": a.file,
	}

	if a.size != nil {
		args["size"] = a.size
	}

	if a.position != nil {
		args["position"] = a.position
	}

	if a.positions != nil {
		args["position"] = a.positions
	}

	return args
}
