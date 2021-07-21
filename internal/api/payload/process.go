package payload

import (
	url2 "github.com/filestack/filestack-go/internal/url"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
)

// ProcessRequest stores request data for Process method.
type ProcessRequest struct {
	ApiKey  string      `json:"apiKey"`
	Sources []string    `json:"sources"`
	Tasks   interface{} `json:"tasks"`
}

// NewProcessRequest constructor.
func NewProcessRequest(
	apiKey string,
	url string,
	storeParams *store.Params,
	security *security.Security,
) ProcessRequest {
	if storeParams == nil {
		storeParams = &store.Params{}
	}

	params := map[string]interface{}{}
	if len(storeParams.FileName) > 0 {
		params["filename"] = storeParams.FileName
	}
	if len(storeParams.WorkFlows) > 0 {
		params["workflows"] = storeParams.WorkFlows
	}
	if len(storeParams.Path) > 0 {
		params["path"] = storeParams.Path
	}
	if len(storeParams.Location) > 0 {
		params["location"] = storeParams.Location
	}
	if len(storeParams.Region) > 0 {
		params["region"] = storeParams.Region
	}
	if len(storeParams.Container) > 0 {
		params["container"] = storeParams.Container
	}
	if len(storeParams.Access) > 0 {
		params["access"] = storeParams.Access
	}

	tasks := []url2.Task{
		{
			Name:   "store",
			Params: params,
		},
	}

	if security != nil {
		tasks = append(tasks, url2.Task{
			Name: "security",
			Params: map[string]interface{}{
				"policy":    security.PolicyB64,
				"signature": security.Signature,
			},
		})
	}

	return ProcessRequest{
		ApiKey:  apiKey,
		Sources: []string{url},
		Tasks:   tasks,
	}
}
