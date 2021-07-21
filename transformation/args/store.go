package args

// StoreArgs arguments for Store transformation.
type StoreArgs struct {
	fileName  *string
	location  *string
	container *string
	path      *string
	region    *string
	access    *string
	base64    bool
	workflows []string
}

// NewStoreArgs constructor.
func NewStoreArgs() *StoreArgs {
	return &StoreArgs{}
}

// SetFileName setter.
func (a *StoreArgs) SetFileName(fileName string) *StoreArgs {
	a.fileName = &fileName
	return a
}

// SetLocation setter.
func (a *StoreArgs) SetLocation(location string) *StoreArgs {
	a.location = &location
	return a
}

// SetContainer setter.
func (a *StoreArgs) SetContainer(container string) *StoreArgs {
	a.container = &container
	return a
}

// SetPath setter.
func (a *StoreArgs) SetPath(path string) *StoreArgs {
	a.path = &path
	return a
}

// SetRegion setter.
func (a *StoreArgs) SetRegion(region string) *StoreArgs {
	a.region = &region
	return a
}

// SetAccess setter.
func (a *StoreArgs) SetAccess(access string) *StoreArgs {
	a.access = &access
	return a
}

// SetBase64 setter.
func (a *StoreArgs) SetBase64(base64 bool) *StoreArgs {
	a.base64 = base64
	return a
}

// SetWorkflows setter.
func (a *StoreArgs) SetWorkflows(workflows []string) *StoreArgs {
	a.workflows = workflows
	return a
}

// ToMap converts this data to a map.
func (a *StoreArgs) ToMap() map[string]interface{} {
	args := map[string]interface{}{}

	if a.fileName != nil {
		args["filename"] = *a.fileName
	}

	if a.location != nil {
		args["location"] = *a.location
	}

	if a.container != nil {
		args["container"] = *a.container
	}

	if a.path != nil {
		args["path"] = *a.path
	}

	if a.region != nil {
		args["region"] = *a.region
	}

	if a.access != nil {
		args["access"] = *a.access
	}

	if a.base64 {
		args["base64"] = true
	}

	if len(a.workflows) > 0 {
		args["workflows"] = a.workflows
	}

	return args
}
