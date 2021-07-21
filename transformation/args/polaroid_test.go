package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPolaroidArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewPolaroidArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("empty map", func(t *testing.T) {

		rotate := 90
		background := "000000"
		color := "FFFFFF"

		mapExpected := map[string]interface{}{
			"rotate":     rotate,
			"background": background,
			"color":      color,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewPolaroidArgs()
		args.SetRotate(rotate)
		args.SetBackground(background)
		args.SetColor(color)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
