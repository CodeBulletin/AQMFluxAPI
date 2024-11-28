--  drop foreing key on sensor
ALTER TABLE Sensor DROP CONSTRAINT fk_device_id;

-- drop device_id from sensor
ALTER TABLE Sensor DROP COLUMN device_id;

-- create a table for sensor_device
CREATE TABLE IF NOT EXISTS SensorDevice (
    sensor_id INT NOT NULL,
    device_id INT NOT NULL,
    PRIMARY KEY (sensor_id, device_id),
    FOREIGN KEY (sensor_id) REFERENCES Sensor(sensor_id),
    FOREIGN KEY (device_id) REFERENCES Device(device_id)
);