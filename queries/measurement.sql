-- name: InsertMeasurement :exec
INSERT INTO measurements (mtime, sensor_id, device_id, attribute_id, mvalue) VALUES ($1, $2, $3, $4, $5);

-- name: GetLatestMeasurement :one
SELECT * FROM measurements WHERE sensor_id = $1 AND device_id = $2 AND attribute_id = $3 ORDER BY mtime DESC LIMIT 1;

-- name: GetHighestMeasurementOfLastHour :one
SELECT MAX(mvalue) FROM measurements WHERE sensor_id = $1 AND device_id = $2 AND attribute_id = $3 AND mtime >= NOW() - INTERVAL '1 hour';
