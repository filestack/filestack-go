package args

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBlurArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewBlurArgs()

		result, _ := json.Marshal(args.ToMap())

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		amount := 2
		mapExpected := map[string]interface{}{
			"amount": amount,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewBlurArgs()
		args.SetAmount(amount)

		result, _ := json.Marshal(args.ToMap())

		assert.Equal(t, string(expected), string(result))
	})

}
