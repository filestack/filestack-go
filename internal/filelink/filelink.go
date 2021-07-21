// Package filelink includes Filelink type.
package filelink

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/filestack/filestack-go/internal/api"
	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/security"
)

// FileLink represents an url
// and allows to operate on the underlying file.
type FileLink struct {
	host           string
	handle         string
	apiKey         string
	security       *security.Security
	requestHandler *api.RequestHandler
}

// New creates a new FileLink or returns an error.
func New(
	host string,
	handle string,
	apiKey string,
	security *security.Security,
	requestHandler *api.RequestHandler,
) (*FileLink, error) {
	if handle == "" {
		return nil, errors.New("cannot create a filelink with an empty handle")
	}

	return &FileLink{
		host:           host,
		handle:         handle,
		apiKey:         apiKey,
		security:       security,
		requestHandler: requestHandler,
	}, nil
}

// MustNew creates a new FileLink and fails execution on error.
func MustNew(
	host string,
	handle string,
	apiKey string,
	security *security.Security,
	requestHandler *api.RequestHandler,
) *FileLink {
	fileLink, err := New(host, handle, apiKey, security, requestHandler)
	if err != nil {
		log.Fatal(err)
	}

	return fileLink
}

// GetHandle gets a handle of the file link.
func (f *FileLink) GetHandle() resource.Handle {
	return resource.NewHandle(f.handle)
}

// AsString generates URL address.
func (f *FileLink) AsString() string {
	URLParts := []string{f.host}
	if f.security != nil {
		URLParts = append(URLParts, f.security.AsString())
	}
	URLParts = append(URLParts, f.handle)

	var URL string
	for i, part := range URLParts {
		if i > 0 {
			URL = URL + "/"
		}
		URL = URL + part
	}

	return URL
}

// Download method allows to get a resource to which file link refers to.
func (f *FileLink) Download(file *os.File) (totalBytes int64, err error) {
	totalBytes = 0
	url := f.AsString()
	response, httpErr := http.Get(url)
	if httpErr != nil {
		err = fmt.Errorf("failed to connect with url %v", url)
		return
	}
	defer response.Body.Close()

	return io.Copy(file, response.Body)
}

// Overwrite method allows to overwrite a resource existing under the related file link.
func (f *FileLink) Overwrite(
	ctx context.Context,
	file io.ReadSeeker,
	base64decode bool,
) error {
	if f.security == nil {
		return fmt.Errorf("security is required to overwrite a filelink")
	}
	_, err := f.requestHandler.Overwrite(ctx, f.handle, file, base64decode, f.security)
	return err
}

// Delete method allows to remove a resource of the related file link.
func (f *FileLink) Delete(ctx context.Context) (err error) {
	if f.security == nil {
		return errors.New("security is required to delete a filelink")
	}

	return f.requestHandler.Delete(ctx, f.handle, f.apiKey, f.security)
}
