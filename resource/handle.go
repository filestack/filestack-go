package resource

import "encoding/base64"

// Handle represents a handle resource.
type Handle string

// NewHandle creates an instance of a handle resource.
func NewHandle(value string) Handle {
	return Handle(value)
}

// AsString coverts Handle to a string.
func (r Handle) AsString() string {
	return string(r)
}

// AsBase64 encodes Handle to base64.
func (r Handle) AsBase64() string {
	return base64.URLEncoding.EncodeToString([]byte(r.AsString()))
}
