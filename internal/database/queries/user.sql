-- name: CreateUser :one
-- Created a new user and returns the full user record
INSERT INTO users (email, first_name, last_name)
VALUES ($1, $2, $3)
RETURNING id, email, first_name, last_name, created_at, updated_at;

-- name: GetUserByID :one
-- Fetches user by primary key (fastest lookup)
SELECT id, email, first_name, last_name, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, first_name, last_name, updated_at, created_at
FROM users
WHERE email = $1;

-- name: GetAllUsers :many
SELECT id, email, first_name, last_name, updated_at, created_at
FROM users
ORDER BY first_name, last_name ASC;

