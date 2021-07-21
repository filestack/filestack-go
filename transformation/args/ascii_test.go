package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewASCIIArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewASCIIArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with params", func(t *testing.T) {

		background := "000000"
		foreground := "ffffff"
		size := 10

		mapExpected := map[string]interface{}{
			"background": background,
			"foreground": foreground,
			"colored":    true,
			"size":       size,
			"reverse":    true,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewASCIIArgs()
		args.SetBackground(background)
		args.SetForeground(foreground)
		args.SetColored()
		args.SetSize(size)
		args.SetReverse()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
