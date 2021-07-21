package config

import (
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
)

// Upload related config parameters.
type Upload struct {
	Intelligent    bool
	StoreParams    *store.Params
	SecurityPolicy *security.Security
}

// NewUploadConfig constructor.
func NewUploadConfig() *Upload {
	return &Upload{}
}
