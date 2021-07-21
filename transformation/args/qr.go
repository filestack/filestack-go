package args

// QR args for QR transformation.
type QR struct {
	version         *int
	errorCorrection *string
	format          *string
}

// NewQR constructor.
func NewQR() *QR {
	return &QR{}
}

// SetVersion setter.
func (a *QR) SetVersion(version int) *QR {
	a.version = &version
	return a
}

// SetErrorCorrection setter.
func (a *QR) SetErrorCorrection(errorCorrection string) *QR {
	a.errorCorrection = &errorCorrection
	return a
}

// SetFormat setter.
func (a *QR) SetFormat(format string) *QR {
	a.format = &format
	return a
}

// ToMap converts this data to a map.
func (a *QR) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.version != nil {
		args["version"] = a.version
	}

	if a.errorCorrection != nil {
		args["error_correction"] = a.errorCorrection
	}

	if a.format != nil {
		args["format"] = a.format
	}

	return args
}
