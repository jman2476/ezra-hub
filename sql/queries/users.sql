-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, phone_number, email, hashed_password)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3, $4)
RETURNING *;

-- name: ClearUsers :exec
DELETE FROM users;

-- name: GetUserByNameEmail :one
SELECT id, created_at, updated_at, name, phone_number, email
FROM users
WHERE name = $1 and email = $2;

-- name: GetUsers :many
SELECT *
FROM users;