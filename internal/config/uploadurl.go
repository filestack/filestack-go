package config

import (
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
)

// UploadURL config parameters.
type UploadURL struct {
	StoreParams    *store.Params
	SecurityPolicy *security.Security
}

// NewUploadURLConfig constructor.
func NewUploadURLConfig() *UploadURL {
	return &UploadURL{}
}
