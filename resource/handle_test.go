package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHandle(t *testing.T) {

	handleValue := "myhandle"
	handle := NewHandle(handleValue)

	t.Run("AsString", func(t *testing.T) {
		assert.Equal(t, handleValue, handle.AsString())
	})

	t.Run("AsBase64", func(t *testing.T) {
		assert.Equal(t, "bXloYW5kbGU=", handle.AsBase64())
	})

}
