package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoundedCornersArgs(t *testing.T) {

	t.Run("radius max", func(t *testing.T) {

		blur := float32(0.3)
		background := "FFFFCC"

		mapExpected := map[string]interface{}{
			"radius":     "max",
			"blur":       blur,
			"background": background,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewRoundedCornersArgs()
		args.SetRadiusMax()
		args.SetBlur(blur)
		args.SetBackground(background)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("radius int", func(t *testing.T) {

		radius := 10
		blur := float32(0.3)
		background := "FFFFCC"

		mapExpected := map[string]interface{}{
			"radius":     radius,
			"blur":       blur,
			"background": background,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewRoundedCornersArgs()
		args.SetRadius(radius)
		args.SetBlur(blur)
		args.SetBackground(background)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
