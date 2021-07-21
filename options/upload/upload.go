// Package upload contains modifiers which apply optional settings to the config.Upload.
package upload

import (
	"errors"

	"github.com/filestack/filestack-go/internal/config"
	"github.com/filestack/filestack-go/options"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
)

// Intelligent enables intelligent upload.
func Intelligent() options.Upload {
	return func(cfg *config.Upload) error {
		cfg.Intelligent = true
		return nil
	}
}

// SecurityPolicy applies a default security policy to the upload methods.
func SecurityPolicy(securityPolicy *security.Security) options.Upload {
	return func(cfg *config.Upload) error {
		if securityPolicy == nil {
			return errors.New("security policy cannot be `nil`")
		}

		cfg.SecurityPolicy = securityPolicy
		return nil
	}
}

// StoreParams applies customized storage parameters.
func StoreParams(storeParams *store.Params) options.Upload {
	return func(cfg *config.Upload) error {
		if storeParams == nil {
			return errors.New("store params cannot be `nil`")
		}

		cfg.StoreParams = storeParams
		return nil
	}
}
