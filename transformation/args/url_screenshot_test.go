package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewURLScreenshotArgs(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		args := NewURLScreenshotArgs()
		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "{}", string(result))
	})

	t.Run("filled up map", func(t *testing.T) {
		args := NewURLScreenshotArgs()
		expectedMap := map[string]interface{}{}

		agent := "mobile"
		args.SetAgent(agent)
		expectedMap["agent"] = agent

		width := 1920
		args.SetWidth(width)
		expectedMap["width"] = width

		height := 1080
		args.SetHeight(height)
		expectedMap["height"] = height

		mode := "window"
		args.SetMode(mode)
		expectedMap["mode"] = mode

		delay := 500
		args.SetDelay(delay)
		expectedMap["delay"] = delay

		orientation := "landscape"
		args.SetOrientation(orientation)
		expectedMap["orientation"] = orientation

		device := "device"
		args.SetDevice(device)
		expectedMap["device"] = device

		result, err := json.Marshal(args.ToMap())
		if err != nil {
			t.Error(err)
		}

		expectedMapJson, _ := json.Marshal(expectedMap)

		assert.Equal(t, string(expectedMapJson), string(result))
	})

}
