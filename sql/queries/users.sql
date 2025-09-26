-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: Reset :exec
TRUNCATE TABLE users CASCADE;

-- name: GetUsers :many
SELECT name FROM users;
