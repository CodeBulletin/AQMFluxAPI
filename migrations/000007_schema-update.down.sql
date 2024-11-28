ALTER TABLE Sensor ADD COLUMN device_id INT NOT NULL;

ALTER TABLE Sensor ADD CONSTRAINT fk_device_id FOREIGN KEY (device_id) REFERENCES Device(device_id);

DROP TABLE IF EXISTS SensorDevice;