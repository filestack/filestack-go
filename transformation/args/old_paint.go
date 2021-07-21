package args

// OldPaintArgs args for Sepia transformation.
type OldPaintArgs struct {
	amount *int
}

// NewOldPaintArgs constructor.
func NewOldPaintArgs() *OldPaintArgs {
	return &OldPaintArgs{}
}

// SetAmount setter.
func (a *OldPaintArgs) SetAmount(amount int) *OldPaintArgs {
	a.amount = &amount
	return a
}

// ToMap converts this data to a map.
func (a *OldPaintArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.amount != nil {
		args["amount"] = a.amount
	}

	return args
}
