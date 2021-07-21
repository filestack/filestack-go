package security

import (
	"encoding/base64"
	"encoding/json"
)

// Policy determines which actions are authorized.
type Policy struct {
	Call      []string `json:"call,omitempty"`
	Container string   `json:"container,omitempty"`
	Expiry    int64    `json:"expiry"`
	Handle    string   `json:"handle,omitempty"`
	MaxSize   int64    `json:"maxSize,omitempty"`
	MinSize   int64    `json:"minSize,omitempty"`
	Path      string   `json:"path,omitempty"`
	URL       string   `json:"url,omitempty"`
}

// AsString represents Policy as a string.
func (p *Policy) AsString() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

// EncodeToB64 encodes policy string to base64 format.
func (p *Policy) EncodeToB64() string {
	return base64.URLEncoding.EncodeToString([]byte(p.AsString()))
}
