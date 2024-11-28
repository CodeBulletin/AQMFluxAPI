-- name: GetSensors :many
SELECT sensor_id, sensor_name, sensor_desc FROM Sensor;

-- name: CreateSensors :exec
INSERT INTO Sensor (sensor_id, sensor_name, sensor_desc) VALUES ($1, $2, $3);

-- name: UpdateSensors :exec
UPDATE Sensor SET sensor_name = $1, sensor_desc = $2 WHERE sensor_id = $3;