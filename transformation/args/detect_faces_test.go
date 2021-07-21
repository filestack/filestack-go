package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDetectFacesArgs(t *testing.T) {

	t.Run("default", func(t *testing.T) {

		args := NewDetectFacesArgs()
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{}
		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with values", func(t *testing.T) {

		maxSize := float32(0.12)
		minSize := float32(0.01)
		color := "FFFFFF"

		args := NewDetectFacesArgs()
		args.SetMaxSize(maxSize)
		args.SetMinSize(minSize)
		args.SetColor(color)
		args.Export()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{
			"maxsize": maxSize,
			"minsize": minSize,
			"color":   color,
			"export":  true,
		}
		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})
}
