-- name: GetUsers :one
SELECT * FROM users;

-- name: GetUserByUsername :one
SELECT * FROM users u
WHERE u.username = $1;

-- name: CreateUser :one
INSERT INTO users(fullname, username, password, permission, login_provider)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;