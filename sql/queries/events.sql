-- name: CreateEvent :one
INSERT INTO events(id, created_at, updated_at, name, owner_id, type, occurs_on, expires_at)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3, $4, $5)
RETURNING *;