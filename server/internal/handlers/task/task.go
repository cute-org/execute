package task

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
)

type Task struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"groupId"`
	CreatorUserID   int       `json:"creatorUserId"`
	CreatorUsername string    `json:"creatorUsername"`
	CreationDate    time.Time `json:"creationDate"`
	DueDate         time.Time `json:"dueDate"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	PointsValue     int       `json:"pointsValue"`
}

type createReq struct {
	DueDate     time.Time `json:"dueDate"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PointsValue int       `json:"pointsValue"`
}

type updateTaskReq struct {
	TaskID      int       `json:"taskId"`
	DueDate     time.Time `json:"dueDate"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PointsValue int       `json:"pointsValue"`
}

// CreateTaskHandler handles POST /task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	groupID, err := auth.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.PointsValue < 0 {
		http.Error(w, "Name required and points must be ≥0", http.StatusBadRequest)
		return
	}

	var taskID int
	err = internal.DB.QueryRow(
		`INSERT INTO tasks
		   (group_id, creator_user_id, due_date, name, description, points_value)
		 VALUES($1,$2,$3,$4,$5,$6) RETURNING id`,
		groupID, userID, req.DueDate, req.Name, req.Description, req.PointsValue,
	).Scan(&taskID)
	if err != nil {
		http.Error(w, "Failed to create task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var username string
	if err := internal.DB.QueryRow(
		"SELECT username FROM users WHERE id=$1", userID,
	).Scan(&username); err != nil {
		http.Error(w, "Failed to retrieve username: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"id":              taskID,
		"creatorUsername": username,
	})
}

// ListTasksHandler handles GET /task
func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	groupID, err := auth.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	rows, err := internal.DB.Query(
		`SELECT
		  t.id,
		  t.group_id,
		  t.creator_user_id,
		  u.username,
		  t.creation_date,
		  t.due_date,
		  t.name,
		  t.description,
		  t.points_value
		FROM tasks t
		JOIN users u ON u.id = t.creator_user_id
		WHERE t.group_id = $1`,
		groupID,
	)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(
			&t.ID,
			&t.GroupID,
			&t.CreatorUserID,
			&t.CreatorUsername,
			&t.CreationDate,
			&t.DueDate,
			&t.Name,
			&t.Description,
			&t.PointsValue,
		); err != nil {
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTaskHandler handles PUT /task
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req updateTaskReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.PointsValue < 0 {
		http.Error(w, "Name required and points must be ≥0", http.StatusBadRequest)
		return
	}

	var creatorID int
	err = internal.DB.QueryRow(
		"SELECT creator_user_id FROM tasks WHERE id=$1", req.TaskID,
	).Scan(&creatorID)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if creatorID != userID {
		http.Error(w, "Forbidden: only the creator can edit", http.StatusForbidden)
		return
	}

	_, err = internal.DB.Exec(
		`UPDATE tasks
		    SET name=$1,
		        description=$2,
		        due_date=$3,
		        points_value=$4
		  WHERE id=$5`,
		req.Name, req.Description, req.DueDate, req.PointsValue, req.TaskID,
	)
	if err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message":"Task updated successfully"}`)
}
