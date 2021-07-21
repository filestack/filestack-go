// Package client contains Client type which is a base service of this SDK.
package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/filestack/filestack-go/internal/api"
	"github.com/filestack/filestack-go/internal/api/payload"
	"github.com/filestack/filestack-go/internal/config"
	"github.com/filestack/filestack-go/internal/filelink"
	"github.com/filestack/filestack-go/internal/upload/intelligent"
	"github.com/filestack/filestack-go/internal/upload/multipart"
	"github.com/filestack/filestack-go/options"
	transformationOptions "github.com/filestack/filestack-go/options/transformation"
	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/transformation"
	"github.com/filestack/filestack-go/transformation/args"
)

// Client simplifies communication and operations related to the FileStack API.
type Client struct {
	apiKey         string
	requestHandler *api.RequestHandler
	config         *config.Client
}

// NewClient creates a new Client instance.
func NewClient(apiKey string, options ...options.Client) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apikey cannot be empty")
	}

	cli := &Client{
		apiKey: apiKey,
		config: config.NewClientConfig(),
	}

	for _, option := range options {
		if err := option(cli.config); err != nil {
			return nil, fmt.Errorf("setting client options has failed: %w", err)
		}
	}

	cli.requestHandler = api.NewRequestHandler(
		cli.config.UploadHost,
		cli.config.APIHost,
		http.DefaultClient,
	)

	return cli, nil
}

// Upload method provides upload functionality for locally stored files.
func (c *Client) Upload(
	ctx context.Context,
	file io.ReadSeeker,
	options ...options.Upload,
) (*filelink.FileLink, error) {
	config := config.NewUploadConfig()
	for _, option := range options {
		if err := option(config); err != nil {
			return nil, fmt.Errorf("setting upload options has failed: %w", err)
		}
	}

	security := c.config.Security
	if config.SecurityPolicy != nil {
		security = config.SecurityPolicy
	}

	if config.Intelligent {
		return intelligent.Upload(
			ctx,
			file,
			c.apiKey,
			c.config.Storage,
			config.StoreParams,
			security,
			c.requestHandler,
			c.config.DefaultFileName,
			c.config.DefaultMimeType,
			c.config.DefaultChunkSize,
			c.config.MinChunkSize,
			c.config.CDNHost,
			c.config.DefaultPartSize,
			c.config.MaxConcurrentUploads,
			c.config.MaxCompleteRetries,
		)
	}

	return multipart.Upload(
		ctx,
		file,
		c.apiKey,
		c.config.Storage,
		config.StoreParams,
		security,
		c.requestHandler,
		c.config.DefaultFileName,
		c.config.DefaultMimeType,
		c.config.DefaultChunkSize,
		c.config.MinChunkSize,
		c.config.CDNHost,
	)
}

// GetApiKey returns the api key that the client has been initialized with.
func (c *Client) GetApiKey() string {
	return c.apiKey
}

// GetSecurity return security policy that the client has been initialized with.
func (c *Client) GetSecurity() *security.Security {
	return c.config.Security
}

// UploadURL method provides upload functionality for externally stored files.
func (c *Client) UploadURL(
	ctx context.Context,
	url string,
	options ...options.UploadURL,
) (*filelink.FileLink, error) {
	config := config.NewUploadURLConfig()
	for _, option := range options {
		if err := option(config); err != nil {
			return nil, fmt.Errorf("failed to apply upload-url options: %w", err)
		}
	}

	security := c.config.Security
	if config.SecurityPolicy != nil {
		security = config.SecurityPolicy
	}

	handle, err := c.requestHandler.Process(ctx, c.config.CDNHost, c.apiKey, url, config.StoreParams, security)
	if err != nil {
		return nil, fmt.Errorf("failed to handle `process` http request: %w", err)
	}
	fileLink, err := filelink.New(c.config.CDNHost, handle, c.apiKey, security, c.requestHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to create a filelink: %w", err)
	}

	return fileLink, nil
}

// ExecuteTransformation performs API call and returns the result as bytes.
func (c *Client) ExecuteTransformation(
	ctx context.Context,
	tr *transformation.Transformation,
) ([]byte, error) {
	return c.requestHandler.GetTransformation(ctx, tr.BuildURL())
}

// NewTransformation creates a new Transformation based on a given resource.
// Returns error on failure.
func (c *Client) NewTransformation(
	res resource.Resource,
	customOptions ...options.Transformation,
) (*transformation.Transformation, error) {
	return transformation.NewTransformation(
		res,
		c.apiKey,
		append(c.getDefaultTransformationOptions(), customOptions...)...,
	)
}

// MustNewTransformation creates a new Transformation based on a given resource.
// Fails execution on error.
func (c *Client) MustNewTransformation(
	res resource.Resource,
	customOptions ...options.Transformation,
) *transformation.Transformation {
	return transformation.MustNewTransformation(
		res,
		c.apiKey,
		append(c.getDefaultTransformationOptions(), customOptions...)...,
	)
}

// NewFileLink getter creates a FileLink based on provided handle string.
// Creates a filelink or returns an error.
func (c *Client) NewFileLink(handle string) (*filelink.FileLink, error) {
	return filelink.New(
		c.config.CDNHost,
		handle,
		c.apiKey,
		c.config.Security,
		c.requestHandler,
	)
}

// MustNewFileLink getter creates a FileLink based on provided handle string.
// Fails execution on error.
func (c *Client) MustNewFileLink(handle string) *filelink.FileLink {
	return filelink.MustNew(
		c.config.CDNHost,
		handle,
		c.apiKey,
		c.config.Security,
		c.requestHandler,
	)
}

// Store allows to persist the result of a transformation.
func (c *Client) Store(
	ctx context.Context,
	transformation *transformation.Transformation,
	storeArgs *args.StoreArgs,
) (responsePayload payload.StoreResponse, err error) {
	url := transformation.Store(storeArgs).BuildURL()
	responsePayload, err = c.requestHandler.Store(ctx, url)
	if err != nil {
		err = fmt.Errorf("sending store request has failed: %w", err)
		return
	}

	return
}

// NewAudioVisual creates a new AudioVisual transformation
// or returns an error.
func (c *Client) NewAudioVisual(
	resource resource.Resource,
	args args.AVConvertOptions,
	customOptions ...options.Transformation,
) (*transformation.AudioVisual, error) {
	tr, err := transformation.NewTransformation(
		resource,
		c.apiKey,
		append(c.getDefaultTransformationOptions(), customOptions...)...,
	)
	if err != nil {
		return nil, err
	}

	return transformation.NewAudioVisual(
		c.config.CDNHost,
		tr.AVConvert(args).BuildURL(),
		c.apiKey,
		c.config.Security,
		c.requestHandler,
	), nil
}

// MustNewAudioVisual creates a new AudioVisual transformation
// or fails execution on error.
func (c *Client) MustNewAudioVisual(
	resource resource.Resource,
	args args.AVConvertOptions,
	customOptions ...options.Transformation,
) *transformation.AudioVisual {
	av, err := c.NewAudioVisual(resource, args, customOptions...)
	if err != nil {
		log.Fatal(err)
	}
	return av
}

// Zip method creates a zip archive from a set of resources.
func (c *Client) Zip(
	ctx context.Context,
	resources ...resource.Resource,
) (responsePayload payload.StoreResponse, err error) {
	zipTransformation, err := transformation.NewTransformationMultiResource(
		resources,
		c.apiKey,
		c.getDefaultTransformationOptions()...,
	)
	if err != nil {
		return
	}
	zipTransformation.Zip()

	return c.Store(ctx, zipTransformation, args.NewStoreArgs())
}

// GetMetadata performs `metadata` transformation task on a given handle.
func (c *Client) GetMetadata(
	ctx context.Context,
	handle resource.Handle,
	customOptions ...options.Transformation,
) (response payload.MetadataResponse, err error) {
	tr, err := c.NewTransformation(
		handle,
		customOptions...,
	)
	if err != nil {
		return
	}

	return c.requestHandler.GetMetadata(ctx, tr.Metadata().BuildURL())
}

// GetTags performs `tags` transformation task on a given handle.
func (c *Client) GetTags(
	ctx context.Context,
	handle resource.Handle,
	customOptions ...options.Transformation,
) (response payload.TagsResponse, err error) {
	tr, err := c.NewTransformation(handle, customOptions...)
	if err != nil {
		return
	}

	return c.requestHandler.GetTags(ctx, tr.Tags().BuildURL())
}

// GetSfw performs `sfw` transformation task on a given handle.
func (c *Client) GetSfw(
	ctx context.Context,
	handle resource.Handle,
	customOptions ...options.Transformation,
) (response payload.SfwResponse, err error) {
	tr, err := c.NewTransformation(handle, customOptions...)
	if err != nil {
		return
	}

	return c.requestHandler.GetSfw(ctx, tr.Sfw().BuildURL())
}

// RunWorkflow starts workflow processing.
// The workflowID can be obtained from
// Filestack Developer Portal: https://dev.filestack.com/
func (c *Client) RunWorkflow(
	ctx context.Context,
	workflowID string,
	rs resource.Resource,
	customOptions ...options.Transformation,
) (*payload.WorkflowRunResponse, error) {
	tr, err := c.NewTransformation(rs, customOptions...)
	if err != nil {
		return nil, err
	}
	if tr.GetSecurity() == nil {
		return nil, errors.New("workflows require security policy")
	}
	tr.RunWorkflow(args.NewRunWorkflowArgs().SetID(workflowID))
	responsePayload, err := c.requestHandler.WorkflowRun(ctx, tr.BuildURL())
	if err != nil {
		return nil, fmt.Errorf("workflow run request has failed: %w", err)
	}
	if len(responsePayload.JobID) == 0 {
		return nil, errors.New("empty response was received")
	}

	return &responsePayload, nil
}

// CheckWorkflowStatus retrieves the status of the provided workflow job.
func (c *Client) CheckWorkflowStatus(
	ctx context.Context,
	jobID string,
	customOptions ...options.Transformation,
) (*payload.WorkflowStatusResponse, error) {
	tr, err := c.NewTransformation(nil, customOptions...)
	if err != nil {
		return nil, err
	}
	if tr.GetSecurity() == nil {
		return nil, errors.New("workflows require security policy")
	}

	tr.WorkflowStatus(args.NewWorkflowStatusArgs().SetJobID(jobID))
	response, err := c.requestHandler.WorkflowStatus(ctx, tr.BuildURL())
	if err != nil {
		return nil, fmt.Errorf("workflow status request has failed: %w", err)
	}
	if len(response.JobID) == 0 {
		return nil, errors.New("empty response was received")
	}

	return &response, nil
}

func (c *Client) getDefaultTransformationOptions() []options.Transformation {
	var defaultOptions []options.Transformation

	if c.config.Security != nil {
		defaultOptions = append(defaultOptions, transformationOptions.SecurityPolicy(c.config.Security))
	}

	return defaultOptions
}
