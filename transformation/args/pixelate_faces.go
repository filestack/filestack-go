package args

// PixelateFacesArgs args for PixelateFaces transformation.
type PixelateFacesArgs struct {
	facesAll   bool
	faces      []int
	maxSize    *float32
	minSize    *float32
	buffer     *int
	amount     *int
	blur       *float32
	filterType *string
}

// NewPixelateFacesArgs constructor.
func NewPixelateFacesArgs() *PixelateFacesArgs {
	return &PixelateFacesArgs{}
}

// SetFacesAll setter.
func (a *PixelateFacesArgs) SetFacesAll() *PixelateFacesArgs {
	a.facesAll = true
	return a
}

// SetFace setter.
func (a *PixelateFacesArgs) SetFace(face int) *PixelateFacesArgs {
	a.faces = []int{face}
	return a
}

// SetFaces setter.
func (a *PixelateFacesArgs) SetFaces(faces []int) *PixelateFacesArgs {
	a.faces = faces
	return a
}

// SetMaxSize setter.
func (a *PixelateFacesArgs) SetMaxSize(maxSize float32) *PixelateFacesArgs {
	a.maxSize = &maxSize
	return a
}

// SetMinSize setter.
func (a *PixelateFacesArgs) SetMinSize(minSize float32) *PixelateFacesArgs {
	a.minSize = &minSize
	return a
}

// SetBuffer setter.
func (a *PixelateFacesArgs) SetBuffer(buffer int) *PixelateFacesArgs {
	a.buffer = &buffer
	return a
}

// SetAmount setter.
func (a *PixelateFacesArgs) SetAmount(amount int) *PixelateFacesArgs {
	a.amount = &amount
	return a
}

// SetBlur setter.
func (a *PixelateFacesArgs) SetBlur(blur float32) *PixelateFacesArgs {
	a.blur = &blur
	return a
}

// SetType setter.
func (a *PixelateFacesArgs) SetType(filterType string) *PixelateFacesArgs {
	a.filterType = &filterType
	return a
}

// ToMap converts this data to a map.
func (a *PixelateFacesArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.facesAll {
		args["faces"] = "all"
	} else {
		if len(a.faces) == 1 {
			args["faces"] = a.faces[0]
		} else if len(a.faces) > 1 {
			args["faces"] = a.faces
		}
	}

	if a.maxSize != nil {
		args["maxsize"] = a.maxSize
	}

	if a.minSize != nil {
		args["minsize"] = a.minSize
	}

	if a.buffer != nil {
		args["buffer"] = a.buffer
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
