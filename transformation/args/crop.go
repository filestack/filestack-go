package args

// CropArgs arguments for Crop transformation.
type CropArgs struct {
	x      int
	y      int
	width  int
	height int
}

// NewCropArgs constructor.
func NewCropArgs(x, y, width, height int) *CropArgs {
	return &CropArgs{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

// ToMap converts this data to a map.
func (a *CropArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{
		"x":      a.x,
		"y":      a.y,
		"width":  a.width,
		"height": a.height,
	}

	return args
}
