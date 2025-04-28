package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
)

// UserProfile represents a user's profile with optional fields omitted when empty
type UserProfile struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name,omitempty"`
	Birthdate   string `json:"birthdate,omitempty"`
	Phone       string `json:"phone,omitempty"`
	GroupID     int64  `json:"group_id,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// UserProfileHandler handles GET requests to fetch the current user profile
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract user ID from auth token
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, "unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Fetch profile in one query and map nulls to defaults
	profile, err := GetUserProfile(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// GetUserProfile retrieves all profile fields at once returning empty or zero values for NULLs
func GetUserProfile(userID int) (UserProfile, error) {
	// Use sql.Null types to scan optional columns
	type raw struct {
		ID          int64          `db:"id"`
		Username    string         `db:"username"`
		DisplayName sql.NullString `db:"display_name"`
		Birthdate   sql.NullTime   `db:"birth_date"`
		Phone       sql.NullString `db:"phone"`
		GroupID     sql.NullInt64  `db:"group_id"`
		CreatedAt   time.Time      `db:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at"`
	}

	query := `
		SELECT id, username, display_name, birth_date, phone, group_id, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	r := raw{}
	err := internal.DB.QueryRow(query, userID).Scan(
		&r.ID,
		&r.Username,
		&r.DisplayName,
		&r.Birthdate,
		&r.Phone,
		&r.GroupID,
		&r.CreatedAt,
		&r.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return UserProfile{}, errors.New("user not found")
	}
	if err != nil {
		return UserProfile{}, err
	}

	// Map raw to UserProfile using empty or zero if NULL
	profile := UserProfile{
		ID:        r.ID,
		Username:  r.Username,
		CreatedAt: r.CreatedAt.Format(time.RFC3339),
		UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
	}
	if r.DisplayName.Valid {
		profile.DisplayName = r.DisplayName.String
	}
	if r.Birthdate.Valid {
		profile.Birthdate = r.Birthdate.Time.Format("2006-01-02")
	}
	if r.Phone.Valid {
		profile.Phone = r.Phone.String
	}
	if r.GroupID.Valid {
		profile.GroupID = r.GroupID.Int64
	}

	return profile, nil
}
