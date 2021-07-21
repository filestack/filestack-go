package args

// PDFConvertArgs args for PDFConvert transformation.
type PDFConvertArgs struct {
	pageOrientation *string
	pageFormat      *string
	pages           string
	metadata        *bool
}

// NewPDFConvertArgs constructor.
func NewPDFConvertArgs() *PDFConvertArgs {
	return &PDFConvertArgs{}
}

// SetPageOrientation setter.
func (a *PDFConvertArgs) SetPageOrientation(pageOrientation string) *PDFConvertArgs {
	a.pageOrientation = &pageOrientation
	return a
}

// SetPageFormat setter.
func (a *PDFConvertArgs) SetPageFormat(pageFormat string) *PDFConvertArgs {
	a.pageFormat = &pageFormat
	return a
}

// SetPages setter.
func (a *PDFConvertArgs) SetPages(pages string) *PDFConvertArgs {
	a.pages = pages
	return a
}

// SetMetadata setter.
func (a *PDFConvertArgs) SetMetadata(metadata bool) *PDFConvertArgs {
	a.metadata = &metadata
	return a
}

// ToMap converts this data to a map.
func (a *PDFConvertArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.pageOrientation != nil {
		args["pageorientation"] = a.pageOrientation
	}

	if a.pageFormat != nil {
		args["pageformat"] = a.pageFormat
	}

	if len(a.pages) > 0 {
		args["pages"] = a.pages
	}

	if a.metadata != nil {
		args["metadata"] = a.metadata
	}

	return args
}
