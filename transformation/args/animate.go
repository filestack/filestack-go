package args

// AnimateArgs args for Animate transformation.
type AnimateArgs struct {
	delay      *int
	loop       *int
	width      *int
	widthMax   bool
	height     *int
	heightMax  bool
	fit        *string
	align      *string
	aligns     []string
	background *string
}

// NewAnimateArgs constructor.
func NewAnimateArgs() *AnimateArgs {
	return &AnimateArgs{}
}

// SetDelay setter.
func (a *AnimateArgs) SetDelay(delay int) *AnimateArgs {
	a.delay = &delay
	return a
}

// SetLoop setter.
func (a *AnimateArgs) SetLoop(loop int) *AnimateArgs {
	a.loop = &loop
	return a
}

// SetWidth setter.
func (a *AnimateArgs) SetWidth(width int) *AnimateArgs {
	a.width = &width
	return a
}

// SetWidthMax setter.
func (a *AnimateArgs) SetWidthMax() *AnimateArgs {
	a.widthMax = true
	return a
}

// SetHeight setter.
func (a *AnimateArgs) SetHeight(height int) *AnimateArgs {
	a.height = &height
	return a
}

// SetHeightMax setter.
func (a *AnimateArgs) SetHeightMax() *AnimateArgs {
	a.heightMax = true
	return a
}

// SetFit setter.
func (a *AnimateArgs) SetFit(fit string) *AnimateArgs {
	a.fit = &fit
	return a
}

// SetAlign setter.
func (a *AnimateArgs) SetAlign(align string) *AnimateArgs {
	a.align = &align
	return a
}

// SetAligns setter.
func (a *AnimateArgs) SetAligns(vertical string, horizontal string) *AnimateArgs {
	a.aligns = []string{vertical, horizontal}
	return a
}

// SetBackground setter.
func (a *AnimateArgs) SetBackground(background string) *AnimateArgs {
	a.background = &background
	return a
}

// ToMap converts this data to a map.
func (a *AnimateArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.delay != nil {
		args["delay"] = a.delay
	}

	if a.loop != nil {
		args["loop"] = a.loop
	}

	if a.widthMax {
		args["width"] = "max"
	} else if a.width != nil {
		args["width"] = a.width
	}

	if a.heightMax {
		args["height"] = "max"
	} else if a.height != nil {
		args["height"] = a.height
	}

	if a.fit != nil {
		args["fit"] = a.fit
	}

	if a.background != nil {
		args["background"] = a.background
	}

	if len(a.aligns) > 0 {
		args["align"] = a.aligns
	} else if a.align != nil {
		args["align"] = a.align
	}

	return args
}
