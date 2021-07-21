package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWatermarkArgs(t *testing.T) {

	t.Run("default", func(t *testing.T) {
		file := "file.jpg"

		args := NewWatermarkArgs(file)
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{
			"file": file,
		}

		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with size and position", func(t *testing.T) {
		file := "file.jpg"
		size := 200
		position := "left"

		args := NewWatermarkArgs(file)
		args.SetSize(size)
		args.SetPosition(position)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{
			"file":     file,
			"size":     size,
			"position": position,
		}

		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with size and positions", func(t *testing.T) {
		file := "file.jpg"
		size := 200
		position := []string{"top", "center"}

		args := NewWatermarkArgs(file)
		args.SetSize(size)
		args.SetPositions(position[0], position[1])
		args.SetPosition("bottom") // this should be ignored

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMap := map[string]interface{}{
			"file":     file,
			"size":     size,
			"position": position,
		}

		expected, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expected), string(result))
	})
}
