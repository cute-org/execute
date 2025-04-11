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
```json
{
  "message": "Session is valid",
  "user": "exampleuser"
}
```

Error Responses:
- `404 Unauthorized/Not Found` — No session token found, or token is invalid/expired.
- `405 Method Not Allowed` — Invalid HTTP method used (only GET is allowed).
- `500 Internal Server Error` — Error retrieving session token.

---
