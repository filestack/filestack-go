package args

// DocToImagesArgs args for DocToImages transformation.
type DocToImagesArgs struct {
	pages        []string
	engine       *string
	format       *string
	quality      *int
	density      *int
	hiddenSlides *bool
}

// NewDocToImagesArgs constructor.
func NewDocToImagesArgs() *DocToImagesArgs {
	return &DocToImagesArgs{}
}

// SetPages setter.
func (a *DocToImagesArgs) SetPages(pages []string) *DocToImagesArgs {
	a.pages = pages
	return a
}

// SetEngine setter.
func (a *DocToImagesArgs) SetEngine(engine string) *DocToImagesArgs {
	a.engine = &engine
	return a
}

// SetFormat setter.
func (a *DocToImagesArgs) SetFormat(format string) *DocToImagesArgs {
	a.format = &format
	return a
}

// SetQuality setter.
func (a *DocToImagesArgs) SetQuality(quality int) *DocToImagesArgs {
	a.quality = &quality
	return a
}

// SetDensity setter.
func (a *DocToImagesArgs) SetDensity(density int) *DocToImagesArgs {
	a.density = &density
	return a
}

// SetHiddenSlides setter.
func (a *DocToImagesArgs) SetHiddenSlides(value bool) *DocToImagesArgs {
	a.hiddenSlides = &value
	return a
}

// ToMap converts this data to a map.
func (a *DocToImagesArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if len(a.pages) > 0 {
		args["pages"] = a.pages
	}

	if a.engine != nil {
		args["engine"] = a.engine
	}

	if a.format != nil {
		args["format"] = a.format
	}

	if a.quality != nil {
		args["quality"] = a.quality
	}

	if a.density != nil {
		args["density"] = a.density
	}

	if a.hiddenSlides != nil {
		args["hidden_slides"] = a.hiddenSlides
	}

	return args
}
