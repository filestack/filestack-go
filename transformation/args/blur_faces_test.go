package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlurFacesArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewBlurFacesArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options - faces all", func(t *testing.T) {

		mapExpected := map[string]interface{}{
			"faces": "all",
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewBlurFacesArgs()
		args.SetFacesAll()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options - one face", func(t *testing.T) {

		face := 2
		mapExpected := map[string]interface{}{
			"faces": face,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewBlurFacesArgs()
		args.SetFace(face)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with options - multiple faces", func(t *testing.T) {

		faces := []int{1, 2, 3}
		maxSize := float32(0.35)
		minSize := float32(0.35)
		buffer := 50
		blur := float32(4)
		blurType := "oval"
		mapExpected := map[string]interface{}{
			"faces":   faces,
			"maxsize": maxSize,
			"minsize": minSize,
			"buffer":  buffer,
			"blur":    blur,
			"type":    blurType,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewBlurFacesArgs()
		args.SetFaces(faces)
		args.SetMaxSize(maxSize)
		args.SetMinSize(minSize)
		args.SetBuffer(buffer)
		args.SetBlur(blur)
		args.SetBlurType(blurType)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
