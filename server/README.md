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
- 400 Bad Request — Missing or invalid input.
- 409 Conflict — Username already exists.

---

### 🔐 POST /login

Login to account.

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

*Error Responses:*
- `401 Unauthorized` — Invalid credentials.
- `400 Bad Request` — Missing fields.

---
