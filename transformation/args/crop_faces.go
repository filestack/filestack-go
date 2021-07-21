package args

// CropFacesArgs arguments for CropFaces transformation.
type CropFacesArgs struct {
	width    *int
	height   *int
	facesAll bool
	faces    []int
	maxSize  *float32
	minSize  *float32
	buffer   *int
	mode     *string
}

// NewCropFacesArgs constructor.
func NewCropFacesArgs() *CropFacesArgs {
	return &CropFacesArgs{}
}

// SetWidth setter.
func (a *CropFacesArgs) SetWidth(width int) *CropFacesArgs {
	a.width = &width
	return a
}

// SetHeight setter.
func (a *CropFacesArgs) SetHeight(height int) *CropFacesArgs {
	a.height = &height
	return a
}

// SetFacesAll setter.
func (a *CropFacesArgs) SetFacesAll() *CropFacesArgs {
	a.facesAll = true
	return a
}

// SetFace setter.
func (a *CropFacesArgs) SetFace(face int) *CropFacesArgs {
	a.faces = []int{face}
	return a
}

// SetFaces setter.
func (a *CropFacesArgs) SetFaces(faces []int) *CropFacesArgs {
	a.faces = faces
	return a
}

// SetMaxSize setter.
func (a *CropFacesArgs) SetMaxSize(maxSize float32) *CropFacesArgs {
	a.maxSize = &maxSize
	return a
}

// SetMinSize setter.
func (a *CropFacesArgs) SetMinSize(minSize float32) *CropFacesArgs {
	a.minSize = &minSize
	return a
}

// SetBuffer setter.
func (a *CropFacesArgs) SetBuffer(buffer int) *CropFacesArgs {
	a.buffer = &buffer
	return a
}

// SetMode setter.
func (a *CropFacesArgs) SetMode(mode string) *CropFacesArgs {
	a.mode = &mode
	return a
}

// ToMap converts this data to a map.
func (a *CropFacesArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.width != nil {
		args["width"] = a.width
	}

	if a.height != nil {
		args["height"] = a.height
	}

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

	if a.mode != nil {
		args["mode"] = a.mode
	}

	return args
}
