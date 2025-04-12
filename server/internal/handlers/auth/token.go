package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// Session represents a user session with an associated expiration time
type Session struct {
	Username  string
	ExpiresAt time.Time
}

const sessionDuration = 7 * 24 * time.Hour

// sessionStore holds the sessions with thread-safe access
var sessionStore = struct {
	sessions map[string]Session
	sync.RWMutex
}{
	sessions: make(map[string]Session),
}

// GenerateSessionToken creates a random, Base64-encoded token to be used as a session identifier
func GenerateSessionToken() (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(tokenBytes), nil
}

// CreateSession adds a new session for the specified username with an expiration timestamp
func CreateSession(username string) (string, error) {
	token, err := GenerateSessionToken()
	if err != nil {
		return "", err
	}
	sessionStore.Lock()
	defer sessionStore.Unlock()
	sessionStore.sessions[token] = Session{
		Username:  username,
		ExpiresAt: time.Now().Add(sessionDuration),
	}
	return token, nil
}

// GetSessionUsername retrieves the username associated with a session token if it is not expired
func GetSessionUsername(token string) (string, bool) {
	sessionStore.RLock()
	session, exists := sessionStore.sessions[token]
	sessionStore.RUnlock()
	if !exists || time.Now().After(session.ExpiresAt) {
		// If the session is expired, clean it up
		DeleteSession(token)
		return "", false
	}
	return session.Username, true
}

// DeleteSession removes a session token from the store
func DeleteSession(token string) {
	sessionStore.Lock()
	defer sessionStore.Unlock()
	delete(sessionStore.sessions, token)
}

// sessionStoreCleanup iterates over the sessions and deletes expired ones
func sessionStoreCleanup() {
	sessionStore.Lock()
	defer sessionStore.Unlock()
	now := time.Now()
	for token, session := range sessionStore.sessions {
		if now.After(session.ExpiresAt) {
			delete(sessionStore.sessions, token)
		}
	}
}

func CleanupExpiredSessions(interval time.Duration) {
	for {
		time.Sleep(interval)
		sessionStoreCleanup()
	}
}
