package args

// PartialBlurArgs args for PartialBlur transformation.
type PartialBlurArgs struct {
	objects    []ImageArea
	amount     *int
	blur       *float32
	filterType *string
}

// NewPartialBlurArgs constructor.
func NewPartialBlurArgs(objects []ImageArea) *PartialBlurArgs {
	return &PartialBlurArgs{
		objects: objects,
	}
}

// SetAmount setter.
func (a *PartialBlurArgs) SetAmount(amount int) *PartialBlurArgs {
	a.amount = &amount
	return a
}

// SetBlur setter.
func (a *PartialBlurArgs) SetBlur(blur float32) *PartialBlurArgs {
	a.blur = &blur
	return a
}

// SetFilterType setter.
func (a *PartialBlurArgs) SetFilterType(filterType string) *PartialBlurArgs {
	a.filterType = &filterType
	return a
}

// ToMap converts this data to a map.
func (a *PartialBlurArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	var objects [][]int
	for _, object := range a.objects {
		objects = append(objects, object.AsArray())
	}
	if len(objects) > 0 {
		args["objects"] = objects
	}

	if a.amount != nil {
		args["amount"] = a.amount
	}

	if a.blur != nil {
		args["blur"] = a.blur
	}

	if a.filterType != nil {
		args["type"] = a.filterType
	}

	return args
}
