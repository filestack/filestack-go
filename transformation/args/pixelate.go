package args

// PixelateArgs args for Sepia transformation.
type PixelateArgs struct {
	amount *int
}

// NewPixelateArgs constructor.
func NewPixelateArgs() *PixelateArgs {
	return &PixelateArgs{}
}

// SetAmount setter.
func (a *PixelateArgs) SetAmount(amount int) *PixelateArgs {
	a.amount = &amount
	return a
}

// ToMap converts this data to a map.
func (a *PixelateArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.amount != nil {
		args["amount"] = a.amount
	}

	return args
}
