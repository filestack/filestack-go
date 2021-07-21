package args

// BlurFacesArgs args for BlurFaces transformation.
type BlurFacesArgs struct {
	facesAll bool
	face     *int
	faces    []int
	maxSize  *float32
	minSize  *float32
	buffer   *int
	amount   *float32
	blur     *float32
	blurType *string
}

// NewBlurFacesArgs constructor.
func NewBlurFacesArgs() *BlurFacesArgs {
	return &BlurFacesArgs{}
}

// SetFacesAll setter.
func (a *BlurFacesArgs) SetFacesAll() *BlurFacesArgs {
	a.facesAll = true
	return a
}

// SetFace setter.
func (a *BlurFacesArgs) SetFace(face int) *BlurFacesArgs {
	a.face = &face
	return a
}

// SetFaces setter.
func (a *BlurFacesArgs) SetFaces(faces []int) *BlurFacesArgs {
	a.faces = faces
	return a
}

// SetMaxSize setter.
func (a *BlurFacesArgs) SetMaxSize(maxSize float32) *BlurFacesArgs {
	a.maxSize = &maxSize
	return a
}

// SetMinSize setter.
func (a *BlurFacesArgs) SetMinSize(minSize float32) *BlurFacesArgs {
	a.minSize = &minSize
	return a
}

// SetBuffer setter.
func (a *BlurFacesArgs) SetBuffer(buffer int) *BlurFacesArgs {
	a.buffer = &buffer
	return a
}

// SetAmount setter.
func (a *BlurFacesArgs) SetAmount(amount float32) *BlurFacesArgs {
	a.amount = &amount
	return a
}

// SetBlur setter.
func (a *BlurFacesArgs) SetBlur(blur float32) *BlurFacesArgs {
	a.blur = &blur
	return a
}

// SetBlurType setter.
func (a *BlurFacesArgs) SetBlurType(blurType string) *BlurFacesArgs {
	a.blurType = &blurType
	return a
}

// ToMap converts this data to a map.
func (a *BlurFacesArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.facesAll {
		args["faces"] = "all"
	} else if len(a.faces) > 0 {
		args["faces"] = a.faces
	} else if a.face != nil {
		args["faces"] = a.face
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

	if a.blurType != nil {
		args["type"] = a.blurType
	}

	return args
}
