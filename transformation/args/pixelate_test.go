package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPixelateArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewPixelateArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		amount := 2
		mapExpected := map[string]interface{}{
			"amount": amount,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewPixelateArgs()
		args.SetAmount(amount)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
