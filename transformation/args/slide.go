package args

// SlideArgs args for Slide transformation.
type SlideArgs struct {
	theme  *string
	engine *string
}

// NewSlideArgs constructor.
func NewSlideArgs() *SlideArgs {
	return &SlideArgs{}
}

// SetTheme setter.
func (a *SlideArgs) SetTheme(theme string) *SlideArgs {
	a.theme = &theme
	return a
}

// SetEngine setter.
func (a *SlideArgs) SetEngine(engine string) *SlideArgs {
	a.engine = &engine
	return a
}

// ToMap converts this data to a map.
func (a *SlideArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.theme != nil {
		args["theme"] = a.theme
	}

	if a.engine != nil {
		args["engine"] = a.engine
	}

	return args
}
