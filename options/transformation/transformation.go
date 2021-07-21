// Package transformation contains modifiers
// which apply optional settings to the config.Transformation.
package transformation

import (
	"errors"

	"github.com/filestack/filestack-go/internal/config"
	"github.com/filestack/filestack-go/options"
	"github.com/filestack/filestack-go/security"
)

// SecurityPolicy applies a default security policy to the Transformation.
func SecurityPolicy(securityPolicy *security.Security) options.Transformation {
	return func(cfg *config.Transformation) error {
		if securityPolicy == nil {
			return errors.New("security policy cannot be `nil`")
		}

		cfg.SecurityPolicy = securityPolicy
		return nil
	}
}

// BaseURL allows to overwrite the default base url used for transformations.
func BaseURL(baseURL string) options.Transformation {
	return func(cfg *config.Transformation) error {
		if len(baseURL) == 0 {
			return errors.New("baseURL cannot be empty")
		}

		cfg.BaseURL = baseURL
		return nil
	}
}
