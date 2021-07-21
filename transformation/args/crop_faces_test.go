package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCropFacesArgs(t *testing.T) {

	t.Run("faces all", func(t *testing.T) {

		width := 10
		height := 20
		minSize := float32(1.1)
		maxSize := float32(100.1)
		mode := "fill"
		buffer := 123

		mapExpected := map[string]interface{}{
			"width":   width,
			"height":  height,
			"maxsize": maxSize,
			"minsize": minSize,
			"mode":    mode,
			"buffer":  buffer,
			"faces":   "all",
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewCropFacesArgs()
		args.SetWidth(width)
		args.SetHeight(height)
		args.SetMinSize(minSize)
		args.SetMaxSize(maxSize)
		args.SetMode(mode)
		args.SetBuffer(buffer)
		args.SetFacesAll()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("faces - one", func(t *testing.T) {

		faces := 5

		mapExpected := map[string]interface{}{
			"faces": faces,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewCropFacesArgs()
		args.SetFace(faces)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("faces - multiple", func(t *testing.T) {

		faces := []int{1, 2, 3}

		mapExpected := map[string]interface{}{
			"faces": faces,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewCropFacesArgs()
		args.SetFaces(faces)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
