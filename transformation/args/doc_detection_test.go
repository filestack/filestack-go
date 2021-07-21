package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocDetection(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewDocDetection()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with params", func(t *testing.T) {

		coords := true
		preprocess := false

		mapExpected := map[string]interface{}{
			"coords":     coords,
			"preprocess": preprocess,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewDocDetection()
		args.SetCoords(coords)
		args.SetPreprocess(preprocess)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
