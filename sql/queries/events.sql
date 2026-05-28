-- name: CreateEvent :one
WITH inserted as (INSERT INTO events(
    id, 
    created_at, updated_at, 
    name, 
    description,
    owner_id, 
    category, 
    occurs_on, expires_at,
    min_volunteers, max_volunteers)
VALUES (gen_random_uuid(), NOW(), NOW(), 
    $1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *)
SELECT inserted.*, u.name as creator_name
FROM inserted
JOIN users u ON u.id = inserted.owner_id;

-- name: AddEventResponder :one
UPDATE events
SET respondants = array_append(respondants, $1),
updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: RemoveEventResponder :one
UPDATE events
SET respondants = array_remove(respondants, $1),
updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: GetEventRespondants :many
SELECT u.id, u.name, u.phone_number, u.email
FROM users u
JOIN events e ON u.id = ANY(e.respondants)
WHERE e.id = $1;

-- name: GetEventsByCategory :many
SELECT * FROM events
WHERE category = ANY($1::GENRE[]);

-- name: GetEventsByOwner :many
SELECT * FROM events
WHERE owner_id = $1;

-- name: UpdateEventByID :one
UPDATE events
SET updated_at = NOW(), name = $1,
description = $2, category = $3, occurs_on = $4,
expires_at = $5, min_volunteers = $6, max_volunteers = $7
WHERE id = $8
RETURNING *;