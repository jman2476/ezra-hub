-- name: CreateUser :one
INSERT INTO users (id, created_at, update_at, name, phone_number, email)
VALUES ($1, NOW(), NOW(), $2, $3, $4)
RETURNING *;