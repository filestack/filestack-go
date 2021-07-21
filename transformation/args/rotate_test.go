package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotateArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		args := NewRotateArgs()
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{}
		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("auto rotate (ignores degrees value)", func(t *testing.T) {
		args := NewRotateArgs()
		args.Auto()
		args.SetDegrees(90)
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{
			"exif": true,
			"deg":  "exif",
		}
		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("by degrees", func(t *testing.T) {
		degrees := 90
		color := "001122"
		args := NewRotateArgs()
		args.SetDegrees(degrees)
		args.SetBackground(color)
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{
			"deg":        degrees,
			"background": color,
		}
		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})

}
