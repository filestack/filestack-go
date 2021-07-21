package resource

import (
	"encoding/base64"
	"fmt"
)

// StorageAlias represents a storage alias resource.
type StorageAlias struct {
	alias string
	path  string
}

// NewStorageAlias creates a new storage alias resource.
func NewStorageAlias(alias string, path string) *StorageAlias {
	return &StorageAlias{
		alias: alias,
		path:  path,
	}
}

// AsString converts a storage alias resource to a string.
func (r *StorageAlias) AsString() string {
	return fmt.Sprintf("src://%s/%s", r.alias, r.path)
}

// AsBase64 converts a storage alias resource to a base64 encoded string.
func (r *StorageAlias) AsBase64() string {
	return base64.URLEncoding.EncodeToString([]byte(r.AsString()))
}
