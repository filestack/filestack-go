package resource

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorageAlias(t *testing.T) {

	alias := "someAlias"
	path := "somePath"

	storageAlias := NewStorageAlias(alias, path)

	t.Run("AsString", func(t *testing.T) {
		fmt.Println(storageAlias.AsString())
		assert.Equal(t, "src://"+alias+"/"+path, storageAlias.AsString())
	})

	t.Run("AsBase64", func(t *testing.T) {
		assert.Equal(t, "c3JjOi8vc29tZUFsaWFzL3NvbWVQYXRo", storageAlias.AsBase64())
	})

}
