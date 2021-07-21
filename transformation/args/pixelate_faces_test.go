package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPixelateFacesArgs(t *testing.T) {

	t.Run("faces all", func(t *testing.T) {

		minSize := float32(1.1)
		maxSize := float32(100.1)
		buffer := 123
		amount := 1
		blur := float32(0.1)
		filterType := "oval"

		mapExpected := map[string]interface{}{
			"maxsize": maxSize,
			"minsize": minSize,
			"buffer":  buffer,
			"faces":   "all",
			"amount":  amount,
			"blur":    blur,
			"type":    filterType,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewPixelateFacesArgs()
		args.SetMinSize(minSize)
		args.SetMaxSize(maxSize)
		args.SetBuffer(buffer)
		args.SetFacesAll()
		args.SetAmount(amount)
		args.SetBlur(blur)
		args.SetType(filterType)

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

		args := NewPixelateFacesArgs()
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

		args := NewPixelateFacesArgs()
		args.SetFaces(faces)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
