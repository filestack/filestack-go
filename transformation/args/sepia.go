package args

// SepiaArgs args for Sepia transformation.
type SepiaArgs struct {
	tone *int
}

// NewSepiaArgs constructor.
func NewSepiaArgs() *SepiaArgs {
	return &SepiaArgs{}
}

// SetTone setter.
func (a *SepiaArgs) SetTone(tone int) *SepiaArgs {
	a.tone = &tone
	return a
}

// ToMap converts this data to a map.
func (a *SepiaArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.tone != nil {
		args["tone"] = a.tone
	}

	return args
}
