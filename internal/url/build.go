package url

import (
	"encoding/base64"
	"encoding/json"

	"github.com/filestack/filestack-go/resource"
)

// Build builds url address from parts.
func Build(
	baseURL string,
	apiKey string,
	security string,
	tasks []Task,
	resources []resource.Resource,
	base64Encode bool,
) (url string) {
	var elements []string
	elements = append(elements, baseURL)

	if len(apiKey) > 0 {
		elements = append(elements, apiKey)
	}

	if len(security) > 0 {
		elements = append(elements, security)
	}

	elements = appendTasks(elements, tasks, base64Encode)

	if resources != nil {
		resourcesElement := ""
		for i, res := range resources {
			if i > 0 {
				resourcesElement += ","
			}
			if base64Encode {
				resourcesElement += "b64://" + res.AsBase64()
			} else {
				resourcesElement += res.AsString()
			}
		}
		if len(resources) > 1 {
			resourcesElement = "[" + resourcesElement + "]"
		}
		elements = append(elements, resourcesElement)
	}

	for i, element := range elements {
		if i > 0 {
			url = url + "/"
		}
		url = url + element
	}

	return
}

func appendTasks(elements []string, tasks []Task, base64Enable bool) []string {
	if !base64Enable {
		for _, task := range tasks {
			elements = append(elements, task.AsString())
		}
		return elements
	}

	elements = append(elements, "b64")

	tasksBytes, _ := json.Marshal(tasks)
	elements = append(elements, base64.URLEncoding.EncodeToString(tasksBytes))

	return elements
}
