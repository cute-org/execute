package user

import (
	"database/sql"
	"errors"
	"time"

	"execute/internal"
)

// GetUserGroupID looks up the group ID associated with a user by their user ID
func GetUserGroupID(userID int) (int64, error) {
	var groupID sql.NullInt64
	err := internal.DB.QueryRow(
		`SELECT group_id FROM users WHERE id = $1`,
		userID,
	).Scan(&groupID)
	if err == sql.ErrNoRows || !groupID.Valid {
		return 0, errors.New("no group associated with user")
	}
	if err != nil {
		return 0, err
	}
	return groupID.Int64, nil
}

// GetUserDisplayName looks up the display name associated with a user by their user ID
func GetUserDisplayName(userID int) (string, error) {
	var displayName sql.NullString
	err := internal.DB.QueryRow(
		`SELECT display_name FROM users WHERE id = $1`,
		userID,
	).Scan(&displayName)
	if err == sql.ErrNoRows || !displayName.Valid {
		return "", errors.New("no display name associated with user")
	}
	if err != nil {
		return "", err
	}
	return displayName.String, nil
}

// GetUserUsername looks up the username associated with a user by their user ID
func GetUserUsername(userID int) (string, error) {
	var username sql.NullString
	err := internal.DB.QueryRow(
		`SELECT username FROM users WHERE id = $1`,
		userID,
	).Scan(&username)
	if err == sql.ErrNoRows || !username.Valid {
		return "", errors.New("no username associated with user")
	}
	if err != nil {
		return "", err
	}
	return username.String, nil
}

// GetUserBirthdate looks up the birth date associated with a user by their user ID
func GetUserBirthdate(userID int) (time.Time, error) {
	var birthdate sql.NullTime
	err := internal.DB.QueryRow(
		`SELECT birth_date FROM users WHERE id = $1`,
		userID,
	).Scan(&birthdate)
	if err == sql.ErrNoRows || !birthdate.Valid {
		return time.Time{}, errors.New("no birthdate associated with user")
	}
	if err != nil {
		return time.Time{}, err
	}
	return birthdate.Time, nil
}

// GetUserPhone looks up the phone number associated with a user by their user ID
func GetUserPhone(userID int) (string, error) {
	var phone sql.NullString
	err := internal.DB.QueryRow(
		`SELECT phone FROM users WHERE id = $1`,
		userID,
	).Scan(&phone)
	if err == sql.ErrNoRows || !phone.Valid {
		return "", errors.New("no phone number associated with user")
	}
	if err != nil {
		return "", err
	}
	return phone.String, nil
}

// GetUserCreatedAt looks up the creation timestamp for a user by their user ID
func GetUserCreatedAt(userID int) (time.Time, error) {
	var createdAt time.Time
	err := internal.DB.QueryRow(
		`SELECT created_at FROM users WHERE id = $1`,
		userID,
	).Scan(&createdAt)
	if err == sql.ErrNoRows {
		return time.Time{}, errors.New("no user found to get created_at")
	}
	if err != nil {
		return time.Time{}, err
	}
	return createdAt, nil
}

// GetUserUpdatedAt looks up the last update timestamp for a user by their user ID
func GetUserUpdatedAt(userID int) (time.Time, error) {
	var updatedAt time.Time
	err := internal.DB.QueryRow(
		`SELECT updated_at FROM users WHERE id = $1`,
		userID,
	).Scan(&updatedAt)
	if err == sql.ErrNoRows {
		return time.Time{}, errors.New("no user found to get updated_at")
	}
	if err != nil {
		return time.Time{}, err
	}
	return updatedAt, nil
}

// GetUserRole looks up the user role by thier user ID
func GetUserRole(userID int) (string, error) {
	var userRole sql.NullString
	err := internal.DB.QueryRow(
		`SELECT role FROM users WHERE id = $1`,
		userID,
	).Scan(&userRole)
	if err == sql.ErrNoRows || !userRole.Valid {
		return "", errors.New("no display name associated with user")
	}
	if err != nil {
		return "", err
	}
	return userRole.String, nil
}
