package args

// BlurArgs args for Blur transformation.
type BlurArgs struct {
	amount *int
}

// NewBlurArgs constructor.
func NewBlurArgs() *BlurArgs {
	return &BlurArgs{}
}

// SetAmount setter.
func (a *BlurArgs) SetAmount(amount int) *BlurArgs {
	a.amount = &amount
	return a
}

// ToMap converts this data to a map.
func (a *BlurArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.amount != nil {
		args["amount"] = a.amount
	}

	return args
}
