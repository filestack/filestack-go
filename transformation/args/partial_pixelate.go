package args

// PartialPixelateArgs args for PartialPixelate transformation.
type PartialPixelateArgs struct {
	objects    []ImageArea
	amount     *int
	blur       *float32
	filterType *string
}

// NewPartialPixelateArgs constructor.
func NewPartialPixelateArgs(objects []ImageArea) *PartialPixelateArgs {
	return &PartialPixelateArgs{
		objects: objects,
	}
}

// SetAmount setter.
func (a *PartialPixelateArgs) SetAmount(amount int) *PartialPixelateArgs {
	a.amount = &amount
	return a
}

// SetBlur setter.
func (a *PartialPixelateArgs) SetBlur(blur float32) *PartialPixelateArgs {
	a.blur = &blur
	return a
}

// SetFilterType setter.
func (a *PartialPixelateArgs) SetFilterType(filterType string) *PartialPixelateArgs {
	a.filterType = &filterType
	return a
}

// ToMap converts this data to a map.
func (a *PartialPixelateArgs) ToMap() map[string]interface{} {
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
