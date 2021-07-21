package args

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCropArgs(t *testing.T) {

	x := 10
	y := 20
	width := 30
	height := 40

	args := NewCropArgs(x, y, width, height)
	result, err := json.Marshal(args.ToMap())
	if err != nil {
		t.Error(err)
	}

	expected, _ := json.Marshal(map[string]interface{}{
		"x":      x,
		"y":      y,
		"width":  width,
		"height": height,
	})

	assert.Equal(t, string(expected), string(result))

}
