package user

import (
	"database/sql"
	"errors"

	"execute/internal"
)

// GetUserGroupID looks up the group ID associated with a user by their user ID
func GetUserGroupID(userID int) (int64, error) {
	var groupID sql.NullInt64
	err := internal.DB.QueryRow(`SELECT group_id FROM users WHERE id = $1`, userID).Scan(&groupID)
	if err == sql.ErrNoRows || !groupID.Valid {
		return 0, errors.New("no group associated with user")
	}
	if err != nil {
		return 0, err
	}
	return groupID.Int64, nil
}
