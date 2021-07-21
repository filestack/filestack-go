package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicy_AsString(t *testing.T) {

	policy := Policy{
		Expiry:    1618989677,
		Call:      []string{"pick", "read"},
		Handle:    "abc",
		URL:       "http://filestack.com",
		Path:      "mypath",
		Container: "s3",
		MinSize:   8,
		MaxSize:   1024,
	}

	policyStr := policy.AsString()
	assert.Equal(t, policyStr, "{\"call\":[\"pick\",\"read\"],\"container\":\"s3\",\"expiry\":1618989677,\"handle\":\"abc\",\"maxSize\":1024,\"minSize\":8,\"path\":\"mypath\",\"url\":\"http://filestack.com\"}")

}

func TestPolicy_EncodeToB64(t *testing.T) {

	policy := Policy{Expiry: 1617228000}
	result := policy.EncodeToB64()

	assert.Equal(t, "eyJleHBpcnkiOjE2MTcyMjgwMDB9", result)

}
