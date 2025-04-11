package internal

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

// Session store for logged in users.
// Maps session token to username.
var sessions = struct {
	m map[string]string
	sync.RWMutex
}{m: make(map[string]string)}

// GenerateSessionToken creates a random, Base64-encoded token to be used as a session identifier.
func GenerateSessionToken() (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(tokenBytes), nil
}
