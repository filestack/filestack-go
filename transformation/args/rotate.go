package args

// RotateArgs arguments for Rotate transformation.
type RotateArgs struct {
	degrees    *int
	exif       bool
	background *string
}

// NewRotateArgs constructor.
func NewRotateArgs() *RotateArgs {
	return &RotateArgs{}
}

// SetDegrees setter.
func (a *RotateArgs) SetDegrees(degrees int) *RotateArgs {
	a.degrees = &degrees
	return a
}

// Auto setter (enables rotation by exif).
func (a *RotateArgs) Auto() *RotateArgs {
	a.exif = true
	return a
}

// SetBackground setter.
func (a *RotateArgs) SetBackground(background string) *RotateArgs {
	a.background = &background
	return a
}

// ToMap converts this data to a map.
func (a *RotateArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.exif {
		args["deg"] = "exif"
		args["exif"] = true
	}

	if a.degrees != nil && !a.exif {
		args["deg"] = *a.degrees
	}

	if a.background != nil {
		args["background"] = *a.background
	}

	return args
}
