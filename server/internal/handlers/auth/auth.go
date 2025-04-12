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

	// Input validation
	if creds.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	if len(creds.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// Fetch salt and password hash from the database
	var salt, storedPasswordHash string
	row := internal.DB.QueryRow("SELECT salt, passwordhash FROM users WHERE username = $1", creds.Username)
	if err := row.Scan(&salt, &storedPasswordHash); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Decode the salt and compute password hash
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		http.Error(w, "Server error decoding salt", http.StatusInternalServerError)
		return
	}
	computedHash := generatePasswordHash(creds.Password, saltBytes)
	computedHashB64 := base64.StdEncoding.EncodeToString(computedHash)

	// Constant time comparison of hashes
	if !compareHashes(storedPasswordHash, computedHashB64) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate session and token
	token, err := CreateSession(creds.Username)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Set the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(sessionDuration),
		// HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TokenResponse{Token: token})
}
