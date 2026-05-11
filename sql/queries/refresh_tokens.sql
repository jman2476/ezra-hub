-- name: CreateRefreshToken :one
INSERT INTO refreshtokens(token, created_at, updated_at, user_id, expires_at)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    date_add(NOW(), '60 Days')
) returning token, created_at, updated_at;

-- name: GetRefreshToken :one
SELECT * FROM refreshtokens
WHERE token = $1;

-- name: GetUserfromRefreshToken :one
SELECT * FROM refreshtokens
JOIN refreshtokens on users.id = refreshtokens.user_id
WHERE refreshtokens.token = $1
AND revoked_at IS NULL
AND expires_at > NOW();

-- name: RevokeToken :exec
UPDATE refreshtokens
SET updated_at = NOW(), revoked_at = NOW()
WHERE token = $1;