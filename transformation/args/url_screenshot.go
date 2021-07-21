package args

// URLScreenshotArgs arguments for URLScreenshot transformation.
type URLScreenshotArgs struct {
	agent       *string
	mode        *string
	width       *int
	height      *int
	delay       *int
	orientation *string
	device      *string
}

// NewURLScreenshotArgs constructor.
func NewURLScreenshotArgs() *URLScreenshotArgs {
	return &URLScreenshotArgs{}
}

// SetAgent setter.
func (a *URLScreenshotArgs) SetAgent(agent string) *URLScreenshotArgs {
	a.agent = &agent
	return a
}

// SetMode setter.
func (a *URLScreenshotArgs) SetMode(mode string) *URLScreenshotArgs {
	a.mode = &mode
	return a
}

// SetWidth setter.
func (a *URLScreenshotArgs) SetWidth(width int) *URLScreenshotArgs {
	a.width = &width
	return a
}

// SetHeight setter.
func (a *URLScreenshotArgs) SetHeight(height int) *URLScreenshotArgs {
	a.height = &height
	return a
}

// SetDelay setter.
func (a *URLScreenshotArgs) SetDelay(delay int) *URLScreenshotArgs {
	a.delay = &delay
	return a
}

// SetOrientation setter.
func (a *URLScreenshotArgs) SetOrientation(orientation string) *URLScreenshotArgs {
	a.orientation = &orientation
	return a
}

// SetDevice setter.
func (a *URLScreenshotArgs) SetDevice(device string) *URLScreenshotArgs {
	a.device = &device
	return a
}

// ToMap converts this data to a map.
func (a *URLScreenshotArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.agent != nil {
		args["agent"] = *a.agent
	}

	if a.width != nil {
		args["width"] = *a.width
	}

	if a.height != nil {
		args["height"] = *a.height
	}

	if a.mode != nil {
		args["mode"] = *a.mode
	}

	if a.delay != nil {
		args["delay"] = *a.delay
	}

	if a.orientation != nil {
		args["orientation"] = *a.orientation
	}

	if a.device != nil {
		args["device"] = *a.device
	}

	return args
}
