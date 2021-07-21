package args

// MinifyJSArgs args for MinifyJS transformation.
type MinifyJSArgs struct {
	gzip             *bool
	useBabelPolyfill *bool
	keepFnName       *bool
	keepClassName    *bool
	mangle           *bool
	mergeVars        *bool
	removeConsole    *bool
	removeUndefined  *bool
	targets          *bool
}

// NewMinifyJSArgs constructor.
func NewMinifyJSArgs() *MinifyJSArgs {
	return &MinifyJSArgs{}
}

// SetGzip setter.
func (a *MinifyJSArgs) SetGzip(gzip bool) *MinifyJSArgs {
	a.gzip = &gzip
	return a
}

// SetUseBabelPolyfill setter.
func (a *MinifyJSArgs) SetUseBabelPolyfill(useBabelPolyfill bool) *MinifyJSArgs {
	a.useBabelPolyfill = &useBabelPolyfill
	return a
}

// SetKeepFnName setter.
func (a *MinifyJSArgs) SetKeepFnName(keepFnName bool) *MinifyJSArgs {
	a.keepFnName = &keepFnName
	return a
}

// SetKeepClassName setter.
func (a *MinifyJSArgs) SetKeepClassName(keepClassName bool) *MinifyJSArgs {
	a.keepClassName = &keepClassName
	return a
}

// SetMangle setter.
func (a *MinifyJSArgs) SetMangle(mangle bool) *MinifyJSArgs {
	a.mangle = &mangle
	return a
}

// SetMergeVars setter.
func (a *MinifyJSArgs) SetMergeVars(mergeVars bool) *MinifyJSArgs {
	a.mergeVars = &mergeVars
	return a
}

// SetRemoveConsole setter.
func (a *MinifyJSArgs) SetRemoveConsole(removeConsole bool) *MinifyJSArgs {
	a.removeConsole = &removeConsole
	return a
}

// SetRemoveUndefined setter.
func (a *MinifyJSArgs) SetRemoveUndefined(removeUndefined bool) *MinifyJSArgs {
	a.removeUndefined = &removeUndefined
	return a
}

// SetTargets setter.
func (a *MinifyJSArgs) SetTargets(targets bool) *MinifyJSArgs {
	a.targets = &targets
	return a
}

// ToMap converts this data to a map.
func (a *MinifyJSArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.gzip != nil {
		args["gzip"] = a.gzip
	}

	if a.useBabelPolyfill != nil {
		args["use_babel_polyfill"] = a.useBabelPolyfill
	}

	if a.keepFnName != nil {
		args["keep_fn_name"] = a.keepFnName
	}

	if a.keepClassName != nil {
		args["keep_class_name"] = a.keepClassName
	}

	if a.mangle != nil {
		args["mangle"] = a.mangle
	}

	if a.mergeVars != nil {
		args["merge_vars"] = a.mergeVars
	}

	if a.removeConsole != nil {
		args["remove_console"] = a.removeConsole
	}

	if a.removeUndefined != nil {
		args["remove_undefined"] = a.removeUndefined
	}

	if a.targets != nil {
		args["targets"] = a.targets
	}

	return args
}
