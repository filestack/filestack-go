package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAnimateArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewAnimateArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		delay := 10
		loop := 5
		width := 100
		height := 200
		fit := "left"
		align := "center"
		background := "00ff00"

		mapExpected := map[string]interface{}{
			"delay":      delay,
			"loop":       loop,
			"width":      width,
			"height":     height,
			"fit":        fit,
			"align":      align,
			"background": background,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewAnimateArgs()
		args.SetDelay(delay)
		args.SetLoop(loop)
		args.SetWidth(width)
		args.SetHeight(height)
		args.SetFit(fit)
		args.SetAlign(align)
		args.SetBackground(background)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		align := "center"
		aligns := []string{"top", "left"}

		mapExpected := map[string]interface{}{
			"align":  aligns,
			"width":  "max",
			"height": "max",
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewAnimateArgs()
		args.SetAlign(align)
		args.SetAligns(aligns[0], aligns[1])
		args.SetWidthMax()
		args.SetWidth(100)
		args.SetHeightMax()
		args.SetHeight(200)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
