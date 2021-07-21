package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTextSentiment(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {

		text := "this sdk is awesome"

		mapExpected := map[string]interface{}{
			"text": text,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewTextSentimentArgs(text)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

	t.Run("with language", func(t *testing.T) {

		text := "this sdk is awesome"
		language := "it"

		mapExpected := map[string]interface{}{
			"text":     text,
			"language": language,
		}
		expected, _ := json.Marshal(mapExpected)

		args := NewTextSentimentArgs(text)
		args.SetLanguage(language)

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(expected), string(result))
	})

}
