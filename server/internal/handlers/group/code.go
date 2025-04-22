package group

import (
	"crypto/rand"
	"encoding/base32"
	"io"
)

// NewCode returns a URL‑safe, base32‑encoded string of length ~8
func NewCode() (string, error) {
	b := make([]byte, 5) // 5 bytes → 8 base32 chars (without padding)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	// base32.StdEncoding with no padding → uppercase A–Z2–7
	code := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b)
	return code, nil
}
