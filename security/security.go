// Package security implements a mechanism which is described here:
// https://www.filestack.com/docs/security/policies/
package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Security is used for requests authentication.
type Security struct {
	PolicyB64 string
	Signature string
}

// NewSecurity creates a new Security based on your secret and policy.
// Both values can be obtained from dev portal.
func NewSecurity(secret string, policy *Policy) *Security {
	policyB64 := policy.EncodeToB64()
	return &Security{
		PolicyB64: policyB64,
		Signature: EncodeSignature(secret, policyB64),
	}
}

// AsString converts Security to a string.
func (s *Security) AsString() string {
	return "security=p:" + s.PolicyB64 + ",s:" + s.Signature
}

// EncodeSignature converts Security to base64 encoded string.
func EncodeSignature(secret string, policyB64 string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(policyB64))
	return hex.EncodeToString(h.Sum(nil))
}
