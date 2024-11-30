-- name: GetMessageslist :many
SELECT id, Title FROM NotificationMessage;

-- name: GetMessages :one
SELECT * FROM NotificationMessage WHERE id = $1;

-- name: CreateMessage :one
INSERT INTO NotificationMessage (Title, Topic, Payload, Tags, MessagePriority) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateMessage :one
UPDATE NotificationMessage SET Title = $1, Topic = $2, Payload = $3, Tags = $4, MessagePriority = $5, updated_at = NOW() WHERE id = $6 RETURNING *;

-- name: DeleteMessage :one
DELETE FROM NotificationMessage WHERE id = $1 RETURNING *;