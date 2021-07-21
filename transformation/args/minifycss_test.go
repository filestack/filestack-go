package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinifyCSSArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewMinifyCSSArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with params", func(t *testing.T) {

		level := 10

		mapExpected := map[string]interface{}{
			"gzip":  true,
			"level": level,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewMinifyCSSArgs()
		args.SetGzip()
		args.SetLevel(level)
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})
}
