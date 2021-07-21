// Package uploadurl contains modifiers which apply optional settings to the config.UploadURL.
package uploadurl

import (
	"errors"

	"github.com/filestack/filestack-go/internal/config"
	"github.com/filestack/filestack-go/options"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
)

// SecurityPolicy allows to overwrite the default security policy applied to the Client.
func SecurityPolicy(securityPolicy *security.Security) options.UploadURL {
	return func(cfg *config.UploadURL) error {
		if securityPolicy == nil {
			return errors.New("security policy cannot be `nil`")
		}

		cfg.SecurityPolicy = securityPolicy
		return nil
	}
}

// StoreParams allows to set custom storage parameters for the upload method.
func StoreParams(storeParams *store.Params) options.UploadURL {
	return func(cfg *config.UploadURL) error {
		if storeParams == nil {
			return errors.New("store params cannot be `nil`")
		}

		cfg.StoreParams = storeParams
		return nil
	}
}
