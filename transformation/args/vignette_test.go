package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVignetteArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewVignetteArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("empty map", func(t *testing.T) {
		amount := 20
		blurMode := "gaussian"
		background := "FF00FF"

		mapExpected := map[string]interface{}{
			"amount":     amount,
			"blurmode":   blurMode,
			"background": background,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewVignetteArgs()
		args.SetAmount(amount)
		args.SetBlurMode(blurMode)
		args.SetBackground(background)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
