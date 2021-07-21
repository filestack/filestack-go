package transformation

import (
	"context"
	"fmt"
	"net/url"

	"github.com/filestack/filestack-go/internal/api"
	"github.com/filestack/filestack-go/internal/filelink"
	"github.com/filestack/filestack-go/security"
)

// AudioVisual allows to perform audio/video conversions.
type AudioVisual struct {
	host           string
	url            string
	apiKey         string
	security       *security.Security
	uuid           string
	timestamp      string
	requestHandler *api.RequestHandler
}

// NewAudioVisual creates a new AudioVisual service.
func NewAudioVisual(
	host string,
	url string,
	apiKey string,
	security *security.Security,
	requestHandler *api.RequestHandler,
) *AudioVisual {
	return &AudioVisual{
		host:           host,
		url:            url,
		apiKey:         apiKey,
		security:       security,
		requestHandler: requestHandler,
	}
}

// ToFileLink converts AudioVisual to FileLink.
func (av *AudioVisual) ToFileLink(ctx context.Context) (fileLink *filelink.FileLink, err error) {
	responsePayload, err := av.requestHandler.AudioVisualInit(ctx, av.url)
	if err != nil {
		err = fmt.Errorf("making av request has failed: %s", err.Error())
		return
	}

	if responsePayload.Status != "completed" {
		err = ConversionNotCompleted
		return
	}

	_, err = url.Parse(responsePayload.Data.URL)
	if err != nil {
		err = ParsingURLFailed
		return
	}

	fileLink, fileLinkErr := filelink.New(av.host, responsePayload.Data.URL, av.apiKey, av.security, av.requestHandler)
	if fileLinkErr != nil {
		err = fmt.Errorf("failed to create a filelink: %w", fileLinkErr)
		return
	}

	return
}

// Status retrieves conversion status by making a http request.
func (av *AudioVisual) Status(ctx context.Context) (status string, err error) {

	responsePayload, err := av.requestHandler.AudioVisualStatus(ctx, av.url)
	if err != nil {
		err = fmt.Errorf("making av status request has failed: %s", err.Error())
		return
	}

	status = responsePayload.Status
	av.timestamp = responsePayload.Timestamp
	av.uuid = responsePayload.UUID

	return
}
