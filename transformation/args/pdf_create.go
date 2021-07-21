package args

// PDFCreateArgs args for PDFCreate transformation.
type PDFCreateArgs struct {
	engine *string
}

// NewPDFCreateArgs constructor.
func NewPDFCreateArgs() *PDFCreateArgs {
	return &PDFCreateArgs{}
}

// SetEngine setter.
func (a *PDFCreateArgs) SetEngine(engine string) *PDFCreateArgs {
	a.engine = &engine
	return a
}

// ToMap converts this data to a map.
func (a *PDFCreateArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.engine != nil {
		args["engine"] = a.engine
	}

	return args
}
