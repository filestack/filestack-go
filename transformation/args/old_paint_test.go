package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOldPaintArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewOldPaintArgs()

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

		args := NewOldPaintArgs()
		args.SetAmount(amount)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
