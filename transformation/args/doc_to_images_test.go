package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocToImagesArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewDocToImagesArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with params", func(t *testing.T) {

		density := 72
		quality := 95
		pages := []string{"1-3", "4-10"}
		format := "png"
		engine := "poppler"

		mapExpected := map[string]interface{}{
			"pages":   pages,
			"density": density,
			"quality": quality,
			"format":  format,
			"engine":  engine,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewDocToImagesArgs()
		args.SetPages(pages)
		args.SetDensity(density)
		args.SetQuality(quality)
		args.SetFormat(format)
		args.SetEngine(engine)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
