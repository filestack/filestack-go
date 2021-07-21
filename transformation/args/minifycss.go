package args

// MinifyCSSArgs args for MinifyCSS transformation.
type MinifyCSSArgs struct {
	level *int
	gzip  bool
}

// NewMinifyCSSArgs constructor.
func NewMinifyCSSArgs() *MinifyCSSArgs {
	return &MinifyCSSArgs{}
}

// SetLevel setter.
func (a *MinifyCSSArgs) SetLevel(level int) *MinifyCSSArgs {
	a.level = &level
	return a
}

// SetGzip setter.
func (a *MinifyCSSArgs) SetGzip() *MinifyCSSArgs {
	a.gzip = true
	return a
}

// ToMap converts this data to a map.
func (a *MinifyCSSArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.level != nil {
		args["level"] = a.level
	}

	if a.gzip {
		args["gzip"] = true
	}

	return args
}
