package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBorderArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewBorderArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with parameters", func(t *testing.T) {

		color := "ff0000"
		background := "00ff00"
		width := 12

		mapExpected := map[string]interface{}{
			"color":      color,
			"background": background,
			"width":      width,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewBorderArgs()
		args.SetColor(color)
		args.SetBackground(background)
		args.SetWidth(width)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
