package config

import (
	"github.com/filestack/filestack-go/security"
)

// Transformation related config parameters.
type Transformation struct {
	SecurityPolicy *security.Security
	BaseURL        string
}

// NewTransformationConfig constructor.
func NewTransformationConfig() *Transformation {
	return &Transformation{
		BaseURL: "https://cdn.filestackcontent.com",
	}
}
