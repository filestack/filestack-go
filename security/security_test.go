package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeSignature(t *testing.T) {

	result := EncodeSignature("MBVEKA6ARFGNDDTVY7IQDZ4HYU", "eyJleHBpcnkiOjE2MTcyMjgwMDB9")
	assert.Equal(t, "8ff61ba46198ad2778933e578e62fdf37ca4e28193655e8cd04f775598165f90", result)

}
