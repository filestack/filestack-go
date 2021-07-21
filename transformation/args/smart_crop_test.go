package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSmartCropArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewSmartCropArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		mode := "auto"
		width := 100
		height := 50
		fillColor := "white"
		coords := true
		mapExpected := map[string]interface{}{
			"mode":       mode,
			"width":      width,
			"height":     height,
			"fill_color": fillColor,
			"coords":     coords,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewSmartCropArgs()
		args.SetMode(mode)
		args.SetWidth(width)
		args.SetHeight(height)
		args.SetFillColor(fillColor)
		args.SetCoords(coords)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
