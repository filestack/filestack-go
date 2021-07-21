package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQR(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		expectedMap := map[string]interface{}{}

		args := NewQR()
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMapJson, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expectedMapJson), string(result))
	})

	t.Run("filled up map", func(t *testing.T) {

		version := 10
		errorCorrection := "H"
		format := "png"

		expectedMap := map[string]interface{}{
			"version":          version,
			"error_correction": errorCorrection,
			"format":           format,
		}

		args := NewQR()
		args.SetVersion(version)
		args.SetErrorCorrection(errorCorrection)
		args.SetFormat(format)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMapJson, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expectedMapJson), string(result))
	})
}
