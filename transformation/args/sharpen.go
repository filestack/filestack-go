package args

// SharpenArgs args for Sharpen transformation.
type SharpenArgs struct {
	amount *int
}

// NewSharpenArgs constructor.
func NewSharpenArgs() *SharpenArgs {
	return &SharpenArgs{}
}

// SetAmount setter.
func (a *SharpenArgs) SetAmount(amount int) *SharpenArgs {
	a.amount = &amount
	return a
}

// ToMap converts this data to a map.
func (a *SharpenArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.amount != nil {
		args["amount"] = a.amount
	}

	return args
}
