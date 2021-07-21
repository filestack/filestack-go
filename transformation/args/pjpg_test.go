package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPJPGArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewPJPGArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options", func(t *testing.T) {

		quality := 50
		metadata := true

		mapExpected := map[string]interface{}{
			"quality":  quality,
			"metadata": metadata,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewPJPGArgs()
		args.SetQuality(quality)
		args.SetMetadata(metadata)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
