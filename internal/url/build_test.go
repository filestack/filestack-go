package url

import (
	"fmt"
	"testing"

	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/transformation/args"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {

	t.Run("external url", func(t *testing.T) {
		baseURL := "http://cdn.filestackcontent.com"
		apiKey := "Ao0CKsWbSr61dG1HAQjtlz"
		tasks := []Task{
			NewTask("vignette", nil),
		}
		extUrl := "https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg"
		res := resource.NewExternalURL(extUrl)

		url := Build(baseURL, apiKey, "", tasks, []resource.Resource{res}, false)

		expected := fmt.Sprintf("%s/%s/vignette/%s", baseURL, apiKey, extUrl)
		assert.Equal(t, expected, url)
	})

	t.Run("external url with base64", func(t *testing.T) {
		baseURL := "http://cdn.filestackcontent.com"
		apiKey := "Ao0CKsWbSr61dG1HAQjtlz"
		tasks := []Task{
			NewTask("vignette", nil),
		}
		res := resource.NewExternalURL("https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg")

		url := Build(baseURL, apiKey, "", tasks, []resource.Resource{res}, true)

		expected := "http://cdn.filestackcontent.com/Ao0CKsWbSr61dG1HAQjtlz/b64/W3siTmFtZSI6InZpZ25ldHRlIiwiUGFyYW1zIjpudWxsfV0=/b64://aHR0cHM6Ly9pLmlwbHNjLmNvbS8xLzAwMDRDWEczOVFYR0k2QjQtQzMyMS1GNC5qcGc="
		assert.Equal(t, expected, url)
	})

	t.Run("external url with security", func(t *testing.T) {
		baseURL := "http://cdn.filestack.com"
		apiKey := "fff8934203527"
		security := "security=p:POLICY,s:SIGNATURE"
		tasks := []Task{
			NewTask("vignette", nil),
		}
		res := resource.NewExternalURL("https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg")

		url := Build(baseURL, apiKey, security, tasks, []resource.Resource{res}, false)

		expected := "http://cdn.filestack.com/fff8934203527/security=p:POLICY,s:SIGNATURE/vignette/https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg"
		assert.Equal(t, expected, url)
	})

	t.Run("external url with security and multiple tasks", func(t *testing.T) {
		baseURL := "http://cdn.filestack.com"
		apiKey := "fff8934203527"
		security := "security=p:POLICY,s:SIGNATURE"
		tasks := []Task{
			NewTask("vignette", nil),
			NewTask("resize", nil),
			NewTask("crop", nil),
		}
		extUrl := "https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg"
		res := resource.NewExternalURL(extUrl)

		url := Build(baseURL, apiKey, security, tasks, []resource.Resource{res}, false)

		expected := fmt.Sprintf("%s/%s/security=p:POLICY,s:SIGNATURE/vignette/resize/crop/%s", baseURL, apiKey, extUrl)
		assert.Equal(t, expected, url)
	})

	t.Run("multiple resources", func(t *testing.T) {
		baseURL := "http://cdn.filestack.com"
		apiKey := "fff8934203527"
		tasks := []Task{
			NewTask("zip", nil),
		}
		externalURL := resource.NewExternalURL("https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg")
		handle1 := resource.NewHandle("handle1")
		handle2 := resource.NewHandle("handle2")

		url := Build(baseURL, apiKey, "", tasks, []resource.Resource{externalURL, handle1, handle2}, false)

		expected := fmt.Sprintf("%s/%s/zip/[https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg,handle1,handle2]", baseURL, apiKey)
		assert.Equal(t, expected, url)
	})

	t.Run("encode tasks in b64", func(t *testing.T) {
		baseURL := "http://cdn.filestackcontent.com"
		apiKey := "Ao0CKsWbSr61dG1HAQjtlz"
		tasks := []Task{
			NewTask("resize", args.NewResizeArgs().SetWidth(100).SetHeight(200).ToMap()),
			NewTask("flip", nil),
			NewTask("border", args.NewBorderArgs().SetWidth(10).ToMap()),
		}
		externalURL := resource.NewExternalURL("https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg")

		url := Build(baseURL, apiKey, "", tasks, []resource.Resource{externalURL}, true)

		expected := "http://cdn.filestackcontent.com/Ao0CKsWbSr61dG1HAQjtlz/b64/W3siTmFtZSI6InJlc2l6ZSIsIlBhcmFtcyI6eyJoZWlnaHQiOjIwMCwid2lkdGgiOjEwMH19LHsiTmFtZSI6ImZsaXAiLCJQYXJhbXMiOm51bGx9LHsiTmFtZSI6ImJvcmRlciIsIlBhcmFtcyI6eyJ3aWR0aCI6MTB9fV0=/b64://aHR0cHM6Ly9pLmlwbHNjLmNvbS8xLzAwMDRDWEczOVFYR0k2QjQtQzMyMS1GNC5qcGc="
		assert.Equal(t, expected, url)
	})

}
