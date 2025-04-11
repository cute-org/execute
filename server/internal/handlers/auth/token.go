package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

// Session store for logged-in users.
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

// GetSessionUsername retrieves the username associated with a session token.
func GetSessionUsername(token string) (string, bool) {
	sessions.RLock()
	defer sessions.RUnlock()
	username, exists := sessions.m[token]
	return username, exists
}

// SetSession associates a session token with a username.
func SetSession(token, username string) {
	sessions.Lock()
	defer sessions.Unlock()
	sessions.m[token] = username
}

// DeleteSession removes a session token from the store.
func DeleteSession(token string) {
	sessions.Lock()
	defer sessions.Unlock()
	delete(sessions.m, token)
}
