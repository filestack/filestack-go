package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPDFConvertArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewPDFConvertArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("empty map", func(t *testing.T) {

		pageOrientation := "portrait"
		pageFormat := "b4"
		pages := "[1,2,3]"
		metadata := true

		mapExpected := map[string]interface{}{
			"pageorientation": pageOrientation,
			"pageformat":      pageFormat,
			"pages":           pages,
			"metadata":        metadata,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewPDFConvertArgs()
		args.SetPageOrientation(pageOrientation)
		args.SetPageFormat(pageFormat)
		args.SetPages(pages)
		args.SetMetadata(metadata)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})
}
