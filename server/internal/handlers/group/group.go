package group

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
	"execute/internal/handlers/user"
)

type createReq struct {
	Name string `json:"name"`
}

type createResp struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type joinReq struct {
	Code string `json:"code"`
}

type resp struct {
	Message string `json:"message"`
}

type updateGroupReq struct {
	Name string `json:"name"`
	Code string `json:"code,omitempty"`
}

type groupInfoResp struct {
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Points      int       `json:"points"`
	PointsScore int       `json:"pointsScore"`
	Meeting     time.Time `json:"meeting,omitempty"`
}

type setMeetingReq struct {
	Time time.Time `json:"time"`
}

// CreateGroupHandler handles POST /group
func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "Group name is required", http.StatusBadRequest)
		return
	}

	code, err := NewCode()
	if err != nil {
		http.Error(w, "Failed to generate code", http.StatusInternalServerError)
		return
	}

	var id int
	err = internal.DB.QueryRow(
		`INSERT INTO groups(name, code, creator_user_id)
         VALUES($1,$2,$3) RETURNING id`,
		req.Name, code, userID,
	).Scan(&id)
	if err != nil {
		http.Error(w, "Could not create group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createResp{ID: id, Code: code})
}

// JoinGroupHandler handles POST /group/join
func JoinGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req joinReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Code == "" {
		http.Error(w, "Join code is required", http.StatusBadRequest)
		return
	}

	var existing sql.NullInt64
	err = internal.DB.QueryRow(
		"SELECT group_id FROM users WHERE id = $1", userID,
	).Scan(&existing)
	if err != nil {
		http.Error(w, "User lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if existing.Valid {
		http.Error(w, "You are already in a group", http.StatusConflict)
		return
	}

	var groupID int
	err = internal.DB.QueryRow(
		"SELECT id FROM groups WHERE code = $1", req.Code,
	).Scan(&groupID)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid join code", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Lookup error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = internal.DB.Exec(
		"UPDATE users SET group_id = $1 WHERE id = $2",
		groupID, userID,
	)
	if err != nil {
		http.Error(w, "Could not join group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp{Message: "Joined group successfully"})
}

// UpdateGroupHandler handles PUT /group
func UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, "user not authenticated", http.StatusUnauthorized)
		return
	}

	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var req updateGroupReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "group name is required", http.StatusBadRequest)
		return
	}

	var query string
	var args []any

	if req.Code != "" {
		query = `UPDATE groups SET name = $1, code = $2 WHERE id = $3 AND creator_user_id = $4`
		args = []any{req.Name, req.Code, groupID, userID}
	} else {
		query = `UPDATE groups SET name = $1 WHERE id = $2 AND creator_user_id = $3`
		args = []any{req.Name, groupID, userID}
	}

	result, err := internal.DB.Exec(query, args...)
	if err != nil {
		if internal.IsUniqueViolation(err) {
			http.Error(w, "group code already in use", http.StatusConflict)
			return
		}
		http.Error(w, "failed to update group", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "no permission to update group or group not found", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "status": "updated", "group_id": %d }`, groupID)
}

// LeaveGroupHandler handles POST /group/leave
func LeaveGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Check that user is actually in a group
	var existing sql.NullInt64
	err = internal.DB.QueryRow(
		"SELECT group_id FROM users WHERE id = $1",
		userID,
	).Scan(&existing)
	if err != nil {
		http.Error(w, "User lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !existing.Valid {
		http.Error(w, "You are not in a group", http.StatusConflict)
		return
	}

	_, err = internal.DB.Exec(
		"UPDATE users SET group_id = NULL WHERE id = $1",
		userID,
	)
	if err != nil {
		http.Error(w, "Could not leave group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp{Message: "Left group successfully"})
}

// GetGroupInfoHandler handles GET /group/info
func GetGroupInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var name, code string
	var points, pointsScore int
	var meeting sql.NullTime
	err = internal.DB.QueryRow(
		`SELECT name, code, points, points_score, meeting FROM groups WHERE id = $1`, groupID,
	).Scan(&name, &code, &points, &pointsScore, &meeting)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := groupInfoResp{
		Name:        name,
		Code:        code,
		Points:      points,
		PointsScore: pointsScore,
	}
	if meeting.Valid {
		resp.Meeting = meeting.Time
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// SetGroupMeetingHandler handles POST /group/meeting
func SetGroupMeetingHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, "user not authenticated", http.StatusUnauthorized)
		return
	}

	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var isCreator bool
	err = internal.DB.QueryRow(
		`SELECT creator_user_id = $1 FROM groups WHERE id = $2`, userID, groupID,
	).Scan(&isCreator)
	if err != nil {
		http.Error(w, "group lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !isCreator {
		http.Error(w, "only group creator can set the meeting", http.StatusForbidden)
		return
	}

	var req setMeetingReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	_, err = internal.DB.Exec(
		`UPDATE groups SET meeting = $1 WHERE id = $2`,
		req.Time, groupID,
	)
	if err != nil {
		http.Error(w, "failed to update meeting time", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp{Message: "Meeting time updated successfully"})
}
