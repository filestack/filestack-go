package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStoreArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		args := NewStoreArgs()
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, "{}", string(result))
	})

	t.Run("filled up map", func(t *testing.T) {
		args := NewStoreArgs()
		expectedMap := map[string]interface{}{}

		filename := "myfile.txt"
		args.SetFileName("myfile.txt")
		expectedMap["filename"] = filename

		location := "s3"
		args.SetLocation(location)
		expectedMap["location"] = location

		container := "container"
		args.SetContainer(container)
		expectedMap["container"] = container

		path := "path"
		args.SetPath(path)
		expectedMap["path"] = path

		region := "us-east-1"
		args.SetRegion(region)
		expectedMap["region"] = region

		access := "private"
		args.SetAccess(access)
		expectedMap["access"] = access

		args.SetBase64(true)
		expectedMap["base64"] = true

		workflows := []string{"4b371df9-508e-49db-a33e-80cd86738d85", "4b371df9-508e-49db-a33e-80cd86738d85"}
		args.SetWorkflows(workflows)
		expectedMap["workflows"] = workflows

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMapJson, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expectedMapJson), string(result))
	})

}
