package args

// ImageArea allows to define a picture area range.
type ImageArea struct {
	x      int
	y      int
	width  int
	height int
}

// NewImageArea constructor.
func NewImageArea(x, y, width, height int) ImageArea {
	return ImageArea{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

// AsArray converts data to an array.
func (i *ImageArea) AsArray() []int {
	return []int{i.x, i.y, i.width, i.height}
}
