-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, phone_number, email, hashed_password)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3, $4)
RETURNING *;

-- name: ClearUsers :exec
DELETE FROM users;

-- name: GetUserByNameEmail :one
SELECT id, created_at, updated_at, name, phone_number, email, subs
FROM users
WHERE name = $1 and email = $2;

-- name: GetUserforLogin :one
SELECT * FROM users
WHERE name = $1 and email = $2;

-- name: GetUsers :many
SELECT *
FROM users;

-- name: GetUserNameOnly :one
SELECT name FROM users
WHERE id = $1;

-- name: SetSubscriptionbyID :exec
UPDATE users
SET subs = $1
WHERE id = $2;

-- name: GetUserSubsbyID :one
SELECT id, subs FROM users
WHERE id = $1;

-- name: UpdateUserbyID :one
UPDATE users
SET name = $1, email = $2, 
phone_number = $3, updated_at = NOW()
WHERE id = $4
RETURNING *;