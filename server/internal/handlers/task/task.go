package task

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
	Step            int       `json:"step"`
	Completed       bool      `json:"completed"`
}

type createReq struct {
	DueDate     time.Time `json:"dueDate"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PointsValue int       `json:"pointsValue"`
	Step        int       `json:"step"`
}

type deleteReq struct {
	TaskID int `json:"taskId"`
}

type updateTaskReq struct {
	TaskID      int       `json:"taskId"`
	DueDate     time.Time `json:"dueDate"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PointsValue int       `json:"pointsValue"`
}

type StepUpdateReq struct {
	TaskID int    `json:"taskId"`
	Action string `json:"action"`
}

type completionReq struct {
	TaskID    int  `json:"taskId"`
	Completed bool `json:"completed"`
}

// CreateTaskHandler handles POST /task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Authenticate user
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Get user's group
	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	// Decode request
	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.PointsValue < 0 {
		http.Error(w, "Name required and points must be ≥0", http.StatusBadRequest)
		return
	}

	// Start transaction
	tx, err := internal.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Lock and check group's point pool
	var poolPoints int
	if err := tx.QueryRow(
		"SELECT points FROM groups WHERE id = $1 FOR UPDATE",
		groupID,
	).Scan(&poolPoints); err != nil {
		http.Error(w, "Failed to fetch points pool: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if poolPoints < req.PointsValue {
		http.Error(w,
			fmt.Sprintf("Not enough points in pool (have %d, need %d)", poolPoints, req.PointsValue),
			http.StatusBadRequest,
		)
		return
	}

	// Deduct points
	if _, err := tx.Exec(
		"UPDATE groups SET points = points - $1 WHERE id = $2",
		req.PointsValue, groupID,
	); err != nil {
		http.Error(w, "Failed to debit points pool: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert task
	var taskID int
	if err := tx.QueryRow(
		`INSERT INTO tasks
		   (group_id, creator_user_id, due_date, name, description, points_value, step)
		 VALUES ($1,$2,$3,$4,$5,$6,$7)
		 RETURNING id`,
		groupID, userID, req.DueDate, req.Name, req.Description, req.PointsValue, req.Step,
	).Scan(&taskID); err != nil {
		http.Error(w, "Failed to create task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch creator username and respond
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

	groupID, err := user.GetUserGroupID(userID)
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
		  t.points_value,
		  t.step,
          t.completed
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
			&t.Step,
			&t.Completed,
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

// UpdateStepHandler handles PATCH /task
func TaskStepHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body
	var req StepUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the action
	if req.Action != "+1" && req.Action != "-1" {
		http.Error(w, "Invalid action. Must be '+1' or '-1'", http.StatusBadRequest)
		return
	}

	// Get the userID from the auth header
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Retrieve the groupID for the user
	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	// Retrieve the task's group ID to ensure the user belongs to the same group
	var taskGroupID, currentStep int
	err = internal.DB.QueryRow(
		"SELECT group_id, step FROM tasks WHERE id=$1", req.TaskID,
	).Scan(&taskGroupID, &currentStep)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Task lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ensure the user is part of the same group
	if taskGroupID != groupID {
		http.Error(w, "Forbidden: You are not in the same group as the task", http.StatusForbidden)
		return
	}

	// Determine the step update
	var stepChange int
	if req.Action == "+1" {
		stepChange = 1
	} else if req.Action == "-1" {
		stepChange = -1
	}

	// Update the step of the task in the database
	_, err = internal.DB.Exec(
		`UPDATE tasks
		 SET step = step + $1
		 WHERE id = $2`,
		stepChange, req.TaskID,
	)
	if err != nil {
		http.Error(w, "Failed to update step: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the updated step value
	err = internal.DB.QueryRow(
		"SELECT step FROM tasks WHERE id=$1", req.TaskID,
	).Scan(&currentStep)
	if err != nil {
		http.Error(w, "Failed to retrieve updated step value: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated task information
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"taskId":  req.TaskID,
		"step":    currentStep,
		"message": "Task step updated successfully",
	})
}

// ToggleTaskCompletionHandler handles PATCH /task/completion
func ToggleTaskCompletionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request
	var req completionReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate user
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Get user's group
	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	// Start transaction
	tx, err := internal.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Lock and fetch group
	var poolPoints int
	if err := tx.QueryRow(
		"SELECT points FROM groups WHERE id = $1 FOR UPDATE",
		groupID,
	).Scan(&poolPoints); err != nil {
		http.Error(w, "Failed to fetch points pool: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Load task details and current completion flag
	var taskGroupID, taskPointsVal int
	var currentCompleted bool
	if err := tx.QueryRow(
		"SELECT group_id, points_value, completed FROM tasks WHERE id=$1 FOR UPDATE",
		req.TaskID,
	).Scan(&taskGroupID, &taskPointsVal, &currentCompleted); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Task lookup failed: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if taskGroupID != groupID {
		http.Error(w, "Forbidden: You are not in the same group as the task", http.StatusForbidden)
		return
	}

	// Prevent duplicate toggles
	if req.Completed && currentCompleted {
		http.Error(w, "Task is already completed", http.StatusBadRequest)
		return
	}
	if !req.Completed && !currentCompleted {
		http.Error(w, "Task is not completed", http.StatusBadRequest)
		return
	}

	// Credit/debit group pool and update group score
	if req.Completed {
		// mark complete: return points & credit group score
		if _, err := tx.Exec(
			"UPDATE groups SET points = points + $1, points_score = points_score + $1 WHERE id = $2",
			taskPointsVal, groupID,
		); err != nil {
			http.Error(w, "Failed to update group points and score: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// undo complete: take points & debit group score
		if poolPoints < taskPointsVal {
			http.Error(w, "Not enough points in pool to undo completion", http.StatusBadRequest)
			return
		}
		if _, err := tx.Exec(
			"UPDATE groups SET points = points - $1, points_score = points_score - $1 WHERE id = $2",
			taskPointsVal, groupID,
		); err != nil {
			http.Error(w, "Failed to update group points and score: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Update task.completed flag
	if _, err := tx.Exec(
		"UPDATE tasks SET completed = $1 WHERE id = $2",
		req.Completed, req.TaskID,
	); err != nil {
		http.Error(w, "Failed to update completion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"taskId":    req.TaskID,
		"completed": req.Completed,
		"message":   "Task completion status updated successfully",
	})
}

// DeleteTaskHandler handles DELETE /task
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Authenticate user
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Decode JSON body
	var req deleteReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Lookup groupID
	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	// Start transaction
	tx, err := internal.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Lock and fetch group pool
	var poolPoints int
	if err := tx.QueryRow(
		"SELECT points FROM groups WHERE id = $1 FOR UPDATE",
		groupID,
	).Scan(&poolPoints); err != nil {
		http.Error(w, "Failed to fetch group points: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Lock and fetch task details
	var taskGroupID, creatorID, pointsVal int
	var completed bool
	err = tx.QueryRow(
		`SELECT group_id, creator_user_id, points_value, completed
		   FROM tasks
		  WHERE id = $1
		    FOR UPDATE`,
		req.TaskID,
	).Scan(&taskGroupID, &creatorID, &pointsVal, &completed)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Permission check: only creator can delete
	if creatorID != userID {
		http.Error(w, "Forbidden: only the creator can delete", http.StatusForbidden)
		return
	}

	// Ensure same group
	if taskGroupID != groupID {
		http.Error(w, "Forbidden: task does not belong to your group", http.StatusForbidden)
		return
	}

	// Return points to pool only if the task is not already completed
	if !completed {
		if _, err := tx.Exec(
			"UPDATE groups SET points = points + $1 WHERE id = $2",
			pointsVal, groupID,
		); err != nil {
			http.Error(w, "Failed to return points to pool: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Delete the task
	if _, err := tx.Exec(
		"DELETE FROM tasks WHERE id = $1",
		req.TaskID,
	); err != nil {
		http.Error(w, "Failed to delete task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"taskId":  req.TaskID,
		"deleted": true,
		"returnedPoints": func() int {
			if !completed {
				return pointsVal
			}
			return 0
		}(),
		"message": fmt.Sprintf("Task %d deleted.%s", req.TaskID,
			func() string {
				if !completed {
					return fmt.Sprintf(" %d points returned to pool.", pointsVal)
				}
				return ""
			}()),
	})
}
