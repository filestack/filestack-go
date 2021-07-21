package args

// FileTypeConversionArgs args for FileTypeConversion transformation.
type FileTypeConversionArgs struct {
	format          *string
	page            *int
	density         *int
	compress        *bool
	qualityInput    bool
	quality         *int
	secure          bool
	docInfo         bool
	strip           bool
	colorSpace      *string
	background      *string
	pageFormat      *string
	pageOrientation *string
}

// NewFileTypeConversionArgs constructor.
func NewFileTypeConversionArgs() *FileTypeConversionArgs {
	return &FileTypeConversionArgs{}
}

// SetFormat setter.
func (a *FileTypeConversionArgs) SetFormat(format string) *FileTypeConversionArgs {
	a.format = &format
	return a
}

// SetPage setter.
func (a *FileTypeConversionArgs) SetPage(page int) *FileTypeConversionArgs {
	a.page = &page
	return a
}

// SetDensity setter.
func (a *FileTypeConversionArgs) SetDensity(density int) *FileTypeConversionArgs {
	a.density = &density
	return a
}

// SetCompress setter.
func (a *FileTypeConversionArgs) SetCompress(compress bool) *FileTypeConversionArgs {
	a.compress = &compress
	return a
}

// SetQualityInput setter.
func (a *FileTypeConversionArgs) SetQualityInput() *FileTypeConversionArgs {
	a.qualityInput = true
	return a
}

// SetQuality setter.
func (a *FileTypeConversionArgs) SetQuality(quality int) *FileTypeConversionArgs {
	a.quality = &quality
	return a
}

// SetSecure setter.
func (a *FileTypeConversionArgs) SetSecure() *FileTypeConversionArgs {
	a.secure = true
	return a
}

// SetDocInfo setter.
func (a *FileTypeConversionArgs) SetDocInfo() *FileTypeConversionArgs {
	a.docInfo = true
	return a
}

// SetStrip setter.
func (a *FileTypeConversionArgs) SetStrip() *FileTypeConversionArgs {
	a.strip = true
	return a
}

// SetColorSpace setter.
func (a *FileTypeConversionArgs) SetColorSpace(colorSpace string) *FileTypeConversionArgs {
	a.colorSpace = &colorSpace
	return a
}

// SetBackground setter.
func (a *FileTypeConversionArgs) SetBackground(background string) *FileTypeConversionArgs {
	a.background = &background
	return a
}

// SetPageFormat setter.
func (a *FileTypeConversionArgs) SetPageFormat(pageFormat string) *FileTypeConversionArgs {
	a.pageFormat = &pageFormat
	return a
}

// SetPageOrientation setter.
func (a *FileTypeConversionArgs) SetPageOrientation(pageOrientation string) *FileTypeConversionArgs {
	a.pageOrientation = &pageOrientation
	return a
}

// ToMap converts this data to a map.
func (a *FileTypeConversionArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.format != nil {
		args["format"] = a.format
	}

	if a.page != nil {
		args["page"] = a.page
	}

	if a.density != nil {
		args["density"] = a.density
	}

	if a.compress != nil {
		args["compress"] = a.compress
	}

	if a.qualityInput {
		args["quality"] = "input"
	} else if a.quality != nil {
		args["quality"] = a.quality
	}

	if a.secure {
		args["secure"] = true
	}

	if a.docInfo {
		args["docinfo"] = true
	}

	if a.strip {
		args["strip"] = true
	}

	if a.colorSpace != nil {
		args["colorspace"] = a.colorSpace
	}

	if a.background != nil {
		args["background"] = a.background
	}

	if a.pageFormat != nil {
		args["pageformat"] = a.pageFormat
	}

	if a.pageOrientation != nil {
		args["pageorientation"] = a.pageOrientation
	}

	return args
}
