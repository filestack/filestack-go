package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCollageArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewCollageArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with one file", func(t *testing.T) {

		fileName := "a.jpg"
		mapExpected := map[string]interface{}{
			"files": []string{fileName},
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewCollageArgs()
		args.SetFile(fileName)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with multiple file", func(t *testing.T) {

		files := []string{"file-one.jpg", "file-two.jpg"}
		margin := 10
		width := 640
		height := 480
		color := "FF00FF"
		fit := "auto"

		mapExpected := map[string]interface{}{
			"files":  files,
			"margin": margin,
			"width":  width,
			"height": height,
			"color":  color,
			"fit":    fit,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewCollageArgs()
		args.SetFiles(files)
		args.SetMargin(margin)
		args.SetWidth(width)
		args.SetHeight(height)
		args.SetColor(color)
		args.SetFit(fit)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})
}
