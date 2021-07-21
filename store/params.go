// Package store contains parameters related to file storage.
package store

// Params allows to define store parameters
// which should overwrite the default settings from the Client.
type Params struct {
	FileName   string
	MimeType   string
	WorkFlows  string
	UploadTags []string
	Path       string
	Location   string
	Region     string
	Container  string
	Access     string
}
