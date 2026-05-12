-- name: CreateOutgoing :one
INSERT INTO outgoingmessages(id, created_at, updated_at, encoded_msg, sent)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, false)
RETURNING *;

-- name: UpdateByIDtoSENT :exec
UPDATE outgoingmessages
SET updated_at=NOW(), sent=true
where id = $1;

-- name: GetMessageByID :one
SELECT * FROM outgoingmessages
WHERE id = $1;
