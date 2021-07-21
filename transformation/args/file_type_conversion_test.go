package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileTypeConversionArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		mapExpected := map[string]interface{}{}
		expected, _ := json.Marshal(mapExpected)

		args := NewFileTypeConversionArgs()

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with params", func(t *testing.T) {

		format := "docx"
		page := 10
		density := 30
		compress := true
		colorSpace := "rgb"
		background := "00ff00"
		pageFormat := "a4"
		pageOrientation := "portrait"

		mapExpected := map[string]interface{}{
			"format":          format,
			"page":            page,
			"density":         density,
			"compress":        compress,
			"quality":         "input",
			"secure":          true,
			"docinfo":         true,
			"strip":           true,
			"colorspace":      colorSpace,
			"background":      background,
			"pageformat":      pageFormat,
			"pageorientation": pageOrientation,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewFileTypeConversionArgs()
		args.SetFormat(format)
		args.SetPage(page)
		args.SetDensity(density)
		args.SetCompress(compress)
		args.SetQualityInput()
		args.SetSecure()
		args.SetDocInfo()
		args.SetStrip()
		args.SetColorSpace(colorSpace)
		args.SetBackground(background)
		args.SetPageFormat(pageFormat)
		args.SetPageOrientation(pageOrientation)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("quality int", func(t *testing.T) {

		quality := 123
		mapExpected := map[string]interface{}{
			"quality": quality,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewFileTypeConversionArgs()
		args.SetQuality(quality)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})
}
