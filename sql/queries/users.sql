-- name: CreateUser :one
INSERT INTO users (id, created_at, update_at, name, phone_number, email)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
RETURNING *;

-- name: ClearUsers :exec
DELETE FROM users;