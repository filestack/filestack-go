package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUpscaleArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewUpscaleArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("empty map", func(t *testing.T) {

		noise := "high"
		upscale := true
		style := "photo"

		mapExpected := map[string]interface{}{
			"noise":   noise,
			"upscale": upscale,
			"style":   style,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewUpscaleArgs()
		args.SetNoise(noise)
		args.SetUpscale(upscale)
		args.SetStyle(style)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
