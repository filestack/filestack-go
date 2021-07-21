package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResizeArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		expectedMap := map[string]interface{}{}
		expectedMapJson, _ := json.Marshal(expectedMap)

		args := NewResizeArgs()
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expectedMapJson), string(result))
	})

	t.Run("filled up map (used both methods SetAlign and SetAligns)", func(t *testing.T) {

		width := 100
		height := 200

		expectedMap := map[string]interface{}{
			"width":  width,
			"height": height,
			"fit":    "max",
			"align":  []string{"top", "left"},
			"filter": "lanczos3",
		}

		args := NewResizeArgs()
		args.SetWidth(width)
		args.SetHeight(height)
		args.SetFit("max")
		args.SetAlign("faces")
		args.SetAligns("top", "left")
		args.SetFilter("lanczos3")

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMapJson, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expectedMapJson), string(result))
	})

	t.Run("filled up map (used only SetAlign method)", func(t *testing.T) {

		width := 100
		height := 200

		expectedMap := map[string]interface{}{
			"width":  width,
			"height": height,
			"fit":    "max",
			"align":  "faces",
			"filter": "lanczos3",
		}

		args := NewResizeArgs()
		args.SetWidth(width)
		args.SetHeight(height)
		args.SetFit("max")
		args.SetAlign("faces")
		args.SetFilter("lanczos3")

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMapJson, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expectedMapJson), string(result))
	})
}
