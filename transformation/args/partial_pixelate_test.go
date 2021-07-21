package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPartialPixelateArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		x := 1
		y := 1
		width := 100
		height := 100

		mapExpected := map[string]interface{}{
			"objects": [][]int{{x, y, width, height}},
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewPartialPixelateArgs([]ImageArea{NewImageArea(x, y, width, height)})

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		x := 1
		y := 1
		width := 100
		height := 100
		amount := 10
		blur := float32(1.5)
		filterType := "oval"

		mapExpected := map[string]interface{}{
			"objects": [][]int{{x, y, width, height}},
			"amount":  amount,
			"blur":    blur,
			"type":    filterType,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewPartialPixelateArgs([]ImageArea{NewImageArea(x, y, width, height)})
		args.SetAmount(amount)
		args.SetBlur(blur)
		args.SetFilterType(filterType)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
