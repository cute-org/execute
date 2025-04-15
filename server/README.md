# Backend API Documentation

## Version: v1

---

### 📝 POST /register

Registers a new user.

**Request Body:**
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

**Request Body:**
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

## 🔒🔧 PUT /user

Updates an existing user's information.
*Request Body:*
```json
{
  "id": 123,
  "username": "newusername",
  "password": "currentPassword123",
  "newpassword": "NewSecurePass456",
  "avatar": "base64encodedImage=="
}
```
*Field Descriptions:*
- `id` (integer) — User ID (required).
- `username` (string) — New username (optional).
- `password` (string) — Current password (required for verification).
- `newpassword` (string) — New password (optional, must be at least 6 characters).
- `avatar` (string) — Base64-encoded avatar image (optional).

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

## 🔒🖼️ GET /avatar

Retrieves a user's avatar in base64 encoded format.

*Query Parameters:*
- `id` (integer) — The user ID (required)

*Success Response:*
- Status: `200 OK`
- Content-Type: `text/plain`
```text

data:image/png;base64,<base64_encoded_data>
```

*Error Responses:*
- `400 Bad Request` — Missing or invalid user ID.
- `405 Method Not Allowed` — Only GET method is allowed.
- `500 Internal Server Error` — Error retrieving avatar data.
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.

*Usage in JS:*
```js
fetch('/avatar?id=1')
  .then(res => res.text())
  .then(data => document.querySelector('img').src = data);
```

---
