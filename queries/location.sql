-- name: GetLocations :many
SELECT location_id, location_name, location_desc FROM DeviceLocation;

-- name: CreateLocation :exec
INSERT INTO DeviceLocation (location_id, location_name, location_desc) VALUES ($1, $2, $3);

-- name: UpdateLocation :exec
UPDATE DeviceLocation SET location_name = $1, location_desc = $2 WHERE location_id = $3;