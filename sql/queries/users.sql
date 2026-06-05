-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, phone_number, email, hashed_password, address)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3, $4, $5)
RETURNING *;

-- name: ClearUsers :exec
DELETE FROM users;

-- name: GetUserByNameEmail :one
SELECT id, created_at, updated_at, name, phone_number, email, address, subs
FROM users
WHERE name = $1 and email = $2;

-- name: GetUserforLogin :one
SELECT * FROM users
WHERE email = $1;

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
phone_number = $3, address = $4,
    updated_at = NOW()
WHERE id = $5
RETURNING *;