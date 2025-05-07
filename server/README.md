# Backend API Documentation

## Version: v1

---

### 📝 POST /register

Registers a new user.

*Request Body:*
```json
{
  "username": "exampleuser",
  "password": "securepassword123"
}
```
*Field Descriptions:*
- `username` (string) — Desired username (must be unique).
- `password` (string) — Must be at least 8 characters, include letters and numbers.

*Success Response:*
- Status: `201 Created`
```json
{
  "message": "User registered successfully",
  "userId": 123
}
```

*Error Responses:*
- `400 Bad Request` — Missing or invalid input.
- `409 Conflict` — Username already exists.
- `405 Method Not Allowed` — Invalid HTTP method used (only POST is allowed).

---

### 🔐 POST /login

Login to account and grant a token with a cookie.

*Request Body:*
```json
{
  "username": "exampleuser",
  "password": "securepassword123"
}
```

*Field Descriptions:*
- `username` (string) — Your registered username.
- `password` (string) — Your account password.

*Success Response:*
- Status: `200 OK`
```json
{
  "token": "Rs/v1EWtzorBIckolXyHmAaMagbj..."
}
```
Cookie: `session_token`

*Error Responses:*
- `401 Unauthorized` — Invalid credentials.
- `400 Bad Request` — Missing fields.
- `405 Method Not Allowed` — Invalid HTTP method used (only POST is allowed).


---

### 🔒🔑 GET /validate

Validates the session token provided in cookies.

*Success Response:*
- Status: `200 OK`
```json
{
  "message": "Session is valid",
  "user": "exampleuser"
}
```

*Error Responses:*
- `405 Method Not Allowed` — Invalid HTTP method used (only GET is allowed).
- `500 Internal Server Error` — Error retrieving session token.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

---

### 🔒📋 GET /user

Fetches all users from the database with their IDs and usernames.

*Success Response:*
- Status: `200 OK`
```json
[
  {
    "id": 1,
    "username": "exampleuser"
  },
  {
    "id": 2,
    "username": "anotheruser"
  }
]
```

*Error Responses:*
- `500 Internal Server Error` — Failed to query users.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

---

### 🔒📄 GET /user/current

Fetches all informations about current user from the database using his token.

*Success Response:*
- Status: `200 OK`
```json
{
  "id": 123,
  "username": "mountain",
  "display_name": "Dew",
  "birthdate": "1985-07-21",
  "phone": "+1234567890",
  "role": "soft drink",
  "group_id": 42,
  "created_at": "2025-02-15T10:34:56Z",
  "updated_at": "2025-04-20T14:12:30Z"
}
```
*Field Descriptions:*
- `id` (integer) — Unique identifier of the user.
- `username` (string) — The user’s login name.
- `display_name` (string, optional) — The user’s chosen display name/profile name.
- `birthdate` (string, optional) — Date of birth in YYYY-MM-DD format.
- `phone` (string, optional) — User’s phone number in international format.
- `role` (string, optional) — The user's role.
- `group_id` (integer, optional) — Identifier for the group the user belongs to.
- `created_at` (string) — ISO-8601 timestamp for when the user was created.
- `updated_at` (string) — ISO-8601 timestamp for the last time the user’s profile was updated.

*Error Responses:*
- `500 Internal Server Error` — Failed to query users.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

---

### 🔒🔧 PUT /user

Updates an existing user's information.
*Request Body:*
```json
{
  "username": "newusername",
  "password": "currentPassword123",
  "newpassword": "NewSecurePass456",
  "avatar": "base64encodedImage==",
  "display_name": "Jane Doe",
  "phone": "+1234567890",
  "birth_date": "1990-05-20",
  "role": "admin"
}
```
*Field Descriptions:*
- `username` (string) — New username (optional).
- `password` (string) — Current password (required for verification).
- `newpassword` (string) — New password (optional, must be at least 6 characters).
- `avatar` (string) — Base64-encoded avatar image (optional).
- `display_name` (string) — Display name shown to other users (optional).
- `phone` (string) — User’s phone number (optional).
- `birth_date` (string) — User’s birth date in YYYY-MM-DD format (optional).
- `role` (string) — User role (e.g. admin, user, etc.) (optional).

*Success Response:*
- Status: `200 OK`
```json
{
  "status": "updated"
}
```

*Error Response:*
- `400 Bad Request` — Missing or invalid input.
- `401 Unauthorized` — Incorrect current password.
- `405 Method Not Allowed` — Only PUT is allowed.
- `413 Request Entity Too Large` — Uploaded file exceeds the size limit.
- `415 Unsupported Media Type` — Content-Type not supported.
- `500 Internal Server Error` — Unexpected error during update.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

---

### 🔒🖼️ GET /avatar

Retrieves a user's avatar in base64 encoded format.

*Query Parameters:*
- `id` (integer) — The user ID (required)

*Success Response:*
- Status: `200 OK`
```json
{
  "avatar": "data:image/png;base64,<base64_encoded_data>"
}
```

*Error Responses:*
- `400 Bad Request` — Missing or invalid user ID.
- `405 Method Not Allowed` — Only GET method is allowed.
- `500 Internal Server Error` — Error retrieving avatar data.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

*Usage in JS:*
```js
fetch('/avatar?id=1')
  .then(res => res.json())
  .then(data => {
    document.querySelector('img').src = data.avatar;
  });
```

---

### 🔒📄 GET /group

Fetches all members of the current user’s group from the database using their session token.

*Success Response:*
- Status: `200 OK`
```json
[
  {
    "id": 456,
    "username": "mountain",
    "role": "member",
    "display_name": "Dew",
    "phone": "+19876543210",
    "birth_date": "1992-03-15"
  },
  {
    "id": 789,
    "username": "carnival",
    "display_name": "Carnival",
    "role": "guest"
  }
]
```
*Field Descriptions:*
- `id` (integer) — Unique identifier of the user.
- `username` (string) — The user’s login name.
- `role` (string) — The user’s role within the group.
- `display_name` (string, optional) — The user’s chosen display/profile name.
- `phone` (string, optional) — User’s phone number in international format.
- `birth_date` (string, optional) — Date of birth in YYYY-MM-DD format.

*Error Responses:*
- `401 Unauthorized` — No session cookie found, or session token is invalid/expired.
- `403 Forbidden` — The user is not associated with any group.
- `500 Internal Server Error` — Failed to query group members or scan database rows.

---

### 🔒👥 POST /group

Creates a new group.

*Request Body:*
```json
{
  "name": "Group Name"
}
```
*Field Descriptions:*
- `name` (string) — Name of the group (required).

*Success Response:*
- Status: `201 Created`
```json
{
  "id": 42,
  "code": "AB12CD"
}
```
*Field Descriptions:*
- `id` (integer) — Unique group ID.
- `code` (string) — Join code for inviting others.

*Error Responses:*
- `400 Bad Request` — Invalid or missing group name.
- `401 Unauthorized` — Not logged in.
- `405 Method Not Allowed` — Only POST is allowed.
- `500 Internal Server Error` — Failed to generate or insert group.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

---

### 🔒👥➕ POST /group/join

Allows a user to join an existing group using a join code. Only users not already in a group can join.

*Request Body:*
```json
{
  "code": "AB12CD"
}
```
*Field Descriptions:*
- `code` (string) — Join code for the group.

*Success Response:*
- Status: `200 OK`
```json
{
  "message": "Joined group successfully"
}
```

*Error Responses:*
- `400 Bad Request` — Missing join code or invalid JSON.
- `401 Unauthorized` — Not logged in.
- `404 Not Found` — Invalid or non-existent group code.
- `405 Method Not` Allowed — Only POST is allowed.
- `409 Conflict` — User is already in a group.
- `500 Internal Server Error` — Database error during join.
- `404 Unauthorized/Not Found` — No session token found, invalid/non-existent group code or token is invalid/expired.

---

### 🔒👥🚪 POST /group/leave

Allows a user to leave their current group. Only users already in a group can leave.

*Success Response:*
- Status: `200 OK`
```json
{
  "message": "Left group successfully"
}
```

*Error Responses:*
- `401 Unauthorized` — Not logged in or session invalid.
- `405 Method Not Allowed` — Only POST is allowed.
- `409 Conflict` — User is not in any group.
- `500 Internal Server Error` — Database error during lookup or update.
- `404 Unauthorized/Not Found` — No session token found, invalid/non-existent group code or token is invalid/expired.

---

### 🔒👥✏️ PUT /group

Allows the creator of a group to update the group's name.

*Request Body:*
```json
{
  "name": "New Group Name"
  "code": "new-group-code"
}
```
*Field Descriptions:*
- `name` (string) — New group name (required).
- `code` (string) — New group code (optional, must be unique).

*Success Response:*
- Status: `200 OK`
```json
{
  "message": "Group updated successfully",
  "group_id": 123
}
```

*Error Responses:*
- `400 Bad Request` — Missing or invalid group name.
- `401 Unauthorized` — Not logged in.
- `403 Forbidden` — User is not the group creator.
- `405 Method Not Allowed` — Only PUT is allowed.
- `500 Internal Server Error` — Failed to update group.
- `404 Unauthorized/Not Found` — No session token found, group not found or token is invalid/expired.

---

### 🔒👥📄 GET /group/info

Retrieves basic information about the authenticated user’s group.

*Success Response:*
- Status: `200 OK`
```json
{
  "name": "Group",
  "code": "XY34ZT",
  "points": 500,
  "pointsScore": 0,
  "meeting": "2025-05-12T18:30:00Z"
}
```
*Field Description:*
- `name` (string) — The group’s display name.
- `code` (string) — The alphanumeric join code for the group.
- `points` (int) — The number of points to use for task creation.
- `pointsScore` (int) — The value of points users gained by completing tasks
- `meeting` (string, optional) — The scheduled meeting time in ISO 8601 format. Only included if a meeting has been set.

*Error Responses:*
- `401 Unauthorized` — No valid session token, or session token is expired/invalid.
- `404 Unauthorized/Not Found` — No session token found, group not found or token is invalid/expired.
- `500 Internal Server Error` — An error occurred while retrieving group information from the database.

---

### 🔒👥 POST /group/meeting

Sets or updates the meeting time for the current user’s group. Only the group creator can update the meeting.

*Request Body:*
```json
{
  "time": "2025-06-01T14:00:00Z"
}
```
*Field Descriptions:*
- `time` (string) — Required. The scheduled meeting time in ISO 8601 format.

*Success Response:*
- Status: `200 OK`
```json
{
  "message": "Meeting time updated successfully"
}
```
*Field Descriptions:*
- `message` (string) — A confirmation message indicating that the meeting time was successfully updated.

*Error Responses:*
- `400 Bad Request` — Invalid request payload (e.g., missing or invalid time).
- `401 Unauthorized` — No valid session token, or the session token is expired/invalid.
- `403 Forbidden` — The authenticated user is not the creator of the group.
- `404 Not Found` — The user is not assigned to any group.
- `500 Internal Server Error` — Failed to update the meeting time in the database.

---

### 🔒📋 POST /task

Creates a new task for a group.

*Request Body:*
```json
{
  "dueDate": "2025-04-20T10:00:00Z",
  "name": "Task Name",
  "description": "Task description",
  "pointsValue": 10
}
```
*Field Descriptions:*
- `dueDate` (string, ISO 8601 date-time) — The due date of the task.
- `name` (string) — The name of the task.
- `description` (string) — A description of the task.
- `pointsValue` (integer) — The points associated with the task (must be ≥ 0).

*Success Response:*
    Status: `201 Created`
```json
{
  "id": 1
}
```

*Error Responses:*
- `400 Bad Request` — Missing or invalid input.
- `401 Unauthorized` — User is not authenticated.
- `403 Forbidden` — User is not a member of the specified group.
- `500 Internal Server Error` — Failed to create task.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

---

### 🔒📋 GET /task

Fetches all tasks for the authenticated user's group.

*Success Response:*
    Status: `200 OK`
```json
[
  {
    "id": 1,
    "groupId": 1,
    "creatorUserId": 123,
    "creatorUsername": "Username"
    "creationDate": "2025-04-15T08:00:00Z",
    "dueDate": "2025-04-20T10:00:00Z",
    "name": "Task Name",
    "description": "Task description",
    "pointsValue": 10
  },
  {
    "id": 2,
    "groupId": 1,
    "creatorUserId": 123,
    "creatorUsername": "Username",
    "creationDate": "2025-04-15T09:00:00Z",
    "dueDate": "2025-04-25T10:00:00Z",
    "name": "Another Task",
    "description": "Another description",
    "pointsValue": 15
  }
]
```

*Error Responses:*
- `401 Unauthorized` — User is not authenticated.
- `403 Forbidden` — User is not a member of any group.
- `500 Internal` Server Error — Failed to fetch tasks.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

---

### 🔒🔄 PUT /task

Updates an existing task.

*Request Body:*
```json
{
  "taskId": 1,
  "dueDate": "2025-04-22T10:00:00Z",
  "name": "Updated Task Name",
  "description": "Updated description",
  "pointsValue": 20
}
```
*Field Descriptions:*
- `taskId` (integer) — The ID of the task to be updated.
- `dueDate` (string, ISO 8601 date-time) — The updated due date of the task.
- `name` (string) — The updated name of the task.
- `description` (string) — The updated description of the task.
- `pointsValue` (integer) — The updated points associated with the task (must be ≥ 0).

*Success Response:*
    Status: `200 OK`
```json
{
  "message": "Task updated successfully"
}
```

*Error Responses:*
- `400 Bad Request` — Missing or invalid input.
- `401 Unauthorized` — User is not authenticated.
- `403 Forbidden` — User is not the creator of the task.
- `500 Internal` Server Error — Failed to update task.
- `404 Unauthorized/Not Found` — No session token found, token is invalid/expired or task not found.

---

### 🔒🏷️ PATCH /task

Updates progress of chosen task.

*Request Body:*
```json
{
  "taskId": 5,
  "action": "+1"
}
```
*Field Descriptions:*
- `taskId` (integer) — The ID of the task whose step is to be updated.
- `action` (string) — The action to perform on the task's step. Possible values are:
    * `"+1"` — Increment the task's step by 1.
    * `"-1"` — Decrement the task's step by 1.

*Success Response:*
- Status: `200 OK`
```json
{
  "taskId": 123,
  "step": 2,
  "message": "Task step updated successfully"
}
```

*Error Responses:*
- `400 Bad Request` — Missing or invalid input, such as an invalid action (e.g., action other than +1 or -1).
- `401 Unauthorized` — User is not authenticated or authorized to perform the action.
- `403 Forbidden` — The user is not part of the same group as the task, or the user is not allowed to modify the step of the task.
- `404 Unauthorized/Not Found` — No session token found, token is invalid/expired or task not found.
- `500 Internal Server Error` — A server error occurred while attempting to update the task's step.

---

### 🔒🏁 PATCH /task/completion

Toggles the completion status of a task within the authenticated user’s group, crediting or debiting the group’s point pool and score.

*Success Response:*
- Status: `200 OK`
```json
{
  "taskId": 5,
  "completed": true,
  "message": "Task completion status updated successfully"
}
```
*Field Description*
- `taskId` (integer) — ID of the task to toggle.
- `completed` (boolean) — `true` to mark as completed, `false` to undo completion.

*Error Responses:*
- `400 Bad Request` — Invalid JSON, missing fields, invalid task ID, duplicate toggle (e.g. completing an already completed task, undoing a non-completed task), or insufficient points in pool to undo.
- `401 Unauthorized` — User not authenticated.
- `403 Forbidden` — User not in same group as the task, or group lookup failed.
- `404 Not Found` — Task not found or invalid/expired session token.
- `405 Method Not Allowed` — HTTP method is not PATCH.
- `500 Internal Server Error` — Database errors (transaction start/commit, query failures, update failures).

---
