package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewModulateArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewModulateArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("empty map", func(t *testing.T) {
		brightness := 10
		saturation := 20
		hue := 30

		mapExpected := map[string]interface{}{
			"brightness": brightness,
			"saturation": saturation,
			"hue":        hue,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewModulateArgs()
		args.SetBrightness(brightness)
		args.SetSaturation(saturation)
		args.SetHue(hue)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})
}
