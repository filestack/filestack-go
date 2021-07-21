package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVConvertOptions_ToMap(t *testing.T) {

	o := AVConvertOptions{
		Width:  800,
		Height: 600,
	}
	m := o.ToMap()
	assert.Equal(t, 2, len(m))
}
