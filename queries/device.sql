-- name: GetAllDeviceInformation :many
SELECT Device.device_id, device_name, Device.location_id, device_desc, IP_addr, MAC_addr, PORT
FROM Device
JOIN DeviceAddr ON Device.device_id = DeviceAddr.device_id
JOIN DeviceLocation ON Device.location_id = DeviceLocation.location_id;

-- name: GetDeviceSensors :many
SELECT Sensor.sensor_id, sensor_name  FROM Sensor 
JOIN SensorDevice ON Sensor.sensor_id = SensorDevice.sensor_id
WHERE SensorDevice.device_id = $1;

-- name: CreateDevice :exec
INSERT INTO Device (device_id, device_name, location_id, device_desc) VALUES ($1, $2, $3, $4);

-- name: CreateDeviceAddr :exec
INSERT INTO DeviceAddr (device_id, IP_addr, MAC_addr, PORT) VALUES ($1, $2, $3, $4);

-- name: AddSensorToDevice :exec
INSERT INTO SensorDevice (sensor_id, device_id) VALUES ($1, $2);

-- name: UpdateDevice :exec
UPDATE Device SET device_name = $2, location_id = $3, device_desc = $4 WHERE device_id = $1;

-- name: UpdateDeviceAddr :exec
UPDATE DeviceAddr SET IP_addr = $2, MAC_addr = $3, PORT = $4 WHERE device_id = $1;

-- name: DeleteDeviceSensors :exec
DELETE FROM SensorDevice WHERE device_id = $1;

-- name: GetDevicesList :many
SELECT device_id, device_name FROM Device;