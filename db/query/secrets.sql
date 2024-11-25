-- name: GetExpiredSecrets :many
SELECT * FROM secrets WHERE expires_at < NOW();

-- name: GetSecretByName :one
SELECT * FROM secrets WHERE name = $1;

-- name: UpdateSecret :exec
UPDATE secrets
SET value = $2, expires_at = $3, updated_at = NOW()
WHERE name = $1;