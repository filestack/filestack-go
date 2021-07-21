package args

// ASCIIArgs args for ASCII transformation.
type ASCIIArgs struct {
	background *string
	foreground *string
	colored    bool
	size       *int
	reverse    bool
}

// NewASCIIArgs constructor.
func NewASCIIArgs() *ASCIIArgs {
	return &ASCIIArgs{}
}

// SetBackground setter.
func (a *ASCIIArgs) SetBackground(background string) *ASCIIArgs {
	a.background = &background
	return a
}

// SetForeground setter.
func (a *ASCIIArgs) SetForeground(foreground string) *ASCIIArgs {
	a.foreground = &foreground
	return a
}

// SetColored setter.
func (a *ASCIIArgs) SetColored() *ASCIIArgs {
	a.colored = true
	return a
}

// SetSize setter.
func (a *ASCIIArgs) SetSize(size int) *ASCIIArgs {
	a.size = &size
	return a
}

// SetReverse setter.
func (a *ASCIIArgs) SetReverse() *ASCIIArgs {
	a.reverse = true
	return a
}

// ToMap converts this data to a map.
func (a *ASCIIArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.background != nil {
		args["background"] = a.background
	}

	if a.foreground != nil {
		args["foreground"] = a.foreground
	}

	if a.colored {
		args["colored"] = true
	}

	if a.size != nil {
		args["size"] = a.size
	}

	if a.reverse {
		args["reverse"] = true
	}

	return args
}
