// Package client contains modifiers which apply optional settings to the config.Client.
package client

import (
	"errors"
	"net/http"

	"github.com/filestack/filestack-go/internal/config"
	"github.com/filestack/filestack-go/options"
	"github.com/filestack/filestack-go/security"
)

// SecurityPolicy applies a default security policy to the client
// that will be used with all the client methods where security policy can be applied.
func SecurityPolicy(securityPolicy *security.Security) options.Client {
	return func(cfg *config.Client) error {
		if securityPolicy == nil {
			return errors.New("security policy cannot be `nil`")
		}

		cfg.Security = securityPolicy
		return nil
	}
}

// MaxConcurrentUploads applies max concurrent upload limit to the client
// that allows to overwrite the default setting used by intelligent upload.
func MaxConcurrentUploads(value int) options.Client {
	return func(cfg *config.Client) error {
		if value == 0 {
			return errors.New("max concurrent uploads cannot be 0")
		}

		cfg.MaxConcurrentUploads = value
		return nil
	}
}

// HTTPClient allows to overwrite the default http client
// with a customized instance.
func HTTPClient(client *http.Client) options.Client {
	return func(cfg *config.Client) error {
		if client == nil {
			return errors.New("http client cannot be `nil`")
		}

		cfg.HTTPClient = client
		return nil
	}
}
