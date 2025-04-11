package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"execute/internal"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
	UserID  *int   `json:"userId,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

// RegisterHandler handles the /register POST endpoint
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if len(creds.Username) == 0 {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	if len(creds.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// Generate a unique 16-byte salt
	saltBytes := make([]byte, 16)
	if _, err := rand.Read(saltBytes); err != nil {
		http.Error(w, "Error generating salt", http.StatusInternalServerError)
		return
	}
	salt := base64.StdEncoding.EncodeToString(saltBytes)

	// Compute Argon2id hash for the password
	hashBytes := generatePasswordHash(creds.Password, saltBytes)
	passwordHash := base64.StdEncoding.EncodeToString(hashBytes)

	// Insert the user into the database and return the new user ID
	var userID int
	err := internal.DB.QueryRow(
		"INSERT INTO users(username, salt, passwordhash) VALUES ($1, $2, $3) RETURNING id",
		creds.Username, salt, passwordHash,
	).Scan(&userID)
	if err != nil {
		if internal.IsUniqueViolation(err) {
			http.Error(w, "Username already exists", http.StatusConflict)
		} else {
			http.Error(w, "User creation failed: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := Response{
		Message: "User registered successfully",
		UserID:  &userID,
	}
	json.NewEncoder(w).Encode(resp)
}

// LoginHandler handles the /login POST endpoint
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate credentials
	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	// Check username and password
	var salt, storedPasswordHash string
	row := internal.DB.QueryRow("SELECT salt, passwordhash FROM users WHERE username = $1", creds.Username)
	if err := row.Scan(&salt, &storedPasswordHash); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Decode the salt
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		http.Error(w, "Error decoding salt", http.StatusInternalServerError)
		return
	}

	// Compute the hash for the password
	computedHash := generatePasswordHash(creds.Password, saltBytes)
	computedHashB64 := base64.StdEncoding.EncodeToString(computedHash)

	// Compare the stored password hash with the computed one
	if !compareHashes(storedPasswordHash, computedHashB64) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate a new session token
	token, err := GenerateSessionToken()
	if err != nil {
		http.Error(w, "Error generating session", http.StatusInternalServerError)
		return
	}

	// Store the session token along with the username
	SetSession(token, creds.Username)

	// Set the session cookie with a 7-day expiration
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// Respond with the session token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TokenResponse{Token: token})
}
