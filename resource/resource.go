// Package resource contains a various types of resources
// like for example handles or files, which are required as
// a source to perform transformation tasks.
package resource

// Resource for transformations
type Resource interface {
	AsString() string
	AsBase64() string
}
