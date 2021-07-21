package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTornEdgesArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewTornEdgesArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with spread", func(t *testing.T) {

		first := 1
		second := 2
		background := "FFFFFF"

		mapExpected := map[string]interface{}{
			"spread":     []int{first, second},
			"background": background,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewTornEdgesArgs()
		args.SetSpread(first, second)
		args.SetBackground(background)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
