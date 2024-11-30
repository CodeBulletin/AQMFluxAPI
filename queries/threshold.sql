-- name: CreateThreshold :one
INSERT INTO Threshold (
    sensor_id,
    device_id,
    attribute_id,
    message_id,
    value1,
    value2,
    frequency,
    operator_id,
    TriggerName,
    TriggerEnabled
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
) RETURNING *;

-- name: UpdateThreshold :one
UPDATE Threshold SET
    sensor_id    = $1,
    device_id    = $2,
    attribute_id = $3,
    message_id   = $4,
    value1       = $5,
    value2       = $6,
    frequency    = $7,
    operator_id  = $8,
    TriggerName  = $9,
    TriggerEnabled = $10
WHERE id = $11 RETURNING *;

-- name: DeleteThreshold :one
DELETE FROM Threshold
WHERE id = $1 RETURNING *;