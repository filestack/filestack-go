package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExternalURL(t *testing.T) {

	url := "https://i.iplsc.com/1/0004CXG39QXGI6B4-C321-F4.jpg"
	externalURL := NewExternalURL(url)

	t.Run("AsString", func(t *testing.T) {
		assert.Equal(t, url, externalURL.AsString())
	})

	t.Run("AsBase64", func(t *testing.T) {
		assert.Equal(t, "aHR0cHM6Ly9pLmlwbHNjLmNvbS8xLzAwMDRDWEczOVFYR0k2QjQtQzMyMS1GNC5qcGc=", externalURL.AsBase64())
	})

}
