package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShadowArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewShadowArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("", func(t *testing.T) {

		color := "FFFFFF"
		background := "000000"
		blur := float32(0.32)
		opacity := 1
		x := 1
		y := 2

		mapExpected := map[string]interface{}{
			"color":      color,
			"background": background,
			"blur":       blur,
			"opacity":    opacity,
			"vector":     []int{x, y},
		}
		expected, err := json.Marshal(mapExpected)
		if err != nil {
			t.Error(err)
		}

		args := NewShadowArgs()
		args.SetColor(color)
		args.SetBackground(background)
		args.SetBlur(blur)
		args.SetOpacity(opacity)
		args.SetVector(x, y)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
