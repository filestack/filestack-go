package resource

import "encoding/base64"

// ExternalURL represents an external url resource.
type ExternalURL struct {
	url     string
	headers map[string]string
}

// NewExternalURL creates an instance of an external url resource.
func NewExternalURL(url string) *ExternalURL {
	return &ExternalURL{
		url: url,
	}
}

// NewExternalURLWithHeaders creates an instance of an external url resource
// and allows to include headers.
func NewExternalURLWithHeaders(url string, headers map[string]string) *ExternalURL {
	return &ExternalURL{
		url:     url,
		headers: headers,
	}
}

// AsString converts an external url resource to a string.
func (r *ExternalURL) AsString() string {
	return r.url
}

// AsBase64 encodes an external url resource to base64.
func (r *ExternalURL) AsBase64() string {
	return base64.URLEncoding.EncodeToString([]byte(r.AsString()))
}
