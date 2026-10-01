package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sign returns the signature header for a payload, so receivers can check
// it came from us.
func Sign(secret, payload []byte) string {
	m := hmac.New(sha256.New, secret)
	m.Write(payload)
	return "sha256=" + hex.EncodeToString(m.Sum(nil))
}
