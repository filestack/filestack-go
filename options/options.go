// Package options defines function types for optional arguments
package options

import "github.com/filestack/filestack-go/internal/config"

// Client represents client options.
type Client func(*config.Client) error

// Upload represents upload options.
type Upload func(*config.Upload) error

// UploadURL represents upload url options.
type UploadURL func(*config.UploadURL) error

// Transformation represents transformation options.
type Transformation func(*config.Transformation) error
