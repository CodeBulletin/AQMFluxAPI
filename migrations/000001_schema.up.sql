CREATE TABLE IF NOT EXISTS Config (
    ckey   VARCHAR(255) PRIMARY KEY,
    cvalue TEXT,
    ctype  INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS DeviceLocation (
    location_id   INT PRIMARY KEY,
    location_name TEXT NOT NULL,
    location_desc TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS  Device (
    device_id   INT PRIMARY KEY,
    device_name TEXT NOT NULL,
    location_id INT NOT NULL,
    device_desc TEXT NOT NULL,
    CONSTRAINT fk_location_id FOREIGN KEY (location_id) REFERENCES devicelocation(location_id)
);

CREATE TABLE IF NOT EXISTS DeviceAddr (
    device_id INT PRIMARY KEY,
    IP_addr   TEXT NOT NULL,
    MAC_addr  TEXT NOT NULL,
    PORT      INT NOT NULL,
    CONSTRAINT fk_device_id FOREIGN KEY (device_id) REFERENCES device(device_id)
);

CREATE TABLE IF NOT EXISTS  DeviceState (
    device_id    INT PRIMARY KEY,
    isRestarting BOOLEAN NOT NULL,
    isRunning    BOOLEAN NOT NULL,
    isUpdating   BOOLEAN NOT NULL,
    CONSTRAINT fk_device_id FOREIGN KEY (device_id) REFERENCES device(device_id)
);

CREATE TABLE IF NOT EXISTS Sensor (
    sensor_id   INT NOT NULL PRIMARY KEY,
    sensor_name TEXT NOT NULL,
    sensor_desc TEXT NOT NULL,
    device_id   INT NOT NULL,
    CONSTRAINT fk_device_id FOREIGN KEY (device_id) REFERENCES device(device_id)
);

CREATE TABLE IF NOT EXISTS Attribute (
   attribute_id   INT PRIMARY KEY,
   attribute_name TEXT NOT NULL,
   attribute_desc TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Measurements (
    mtime        TIMESTAMPTZ      NOT NULL,
    sensor_id    INT              NOT NULL,
    device_id    INT              NOT NULL,
    attribute_id INT              NOT NULL,
    mvalue       DOUBLE PRECISION NOT NULL,
    CONSTRAINT fk_device_id    FOREIGN KEY (device_id)    REFERENCES device(device_id),
    CONSTRAINT fk_sensor_id    FOREIGN KEY (sensor_id)    REFERENCES sensor(sensor_id),
    CONSTRAINT fk_attribute_id FOREIGN KEY (attribute_id) REFERENCES attribute(attribute_id),
    PRIMARY KEY (mtime, sensor_id, device_id, attribute_id)
);

CREATE TABLE IF NOT EXISTS Operator (
    id        SERIAL PRIMARY KEY,   
    op        VARCHAR(10) NOT NULL,
    variables INT NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS NotificationMessage (
    id              SERIAL PRIMARY KEY,
    Title           VARCHAR(255) NOT NULL,
    Topic           VARCHAR(255) NOT NULL,
    Payload         TEXT NOT NULL,
    Tags            TEXT,
    MessagePriority INT NOT NULL DEFAULT 1,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Threshold (
    id             SERIAL PRIMARY KEY,
    sensor_id      INT NOT NULL,
    device_id      INT NOT NULL,
    attribute_id   INT NOT NULL,
    message_id     INT NOT NULL,
    value1         FLOAT NOT NULL,
    value2         FLOAT,
    frequency      INT NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_triggered TIMESTAMP,
    operator_id    INT NOT NULL,
    TriggerName    VARCHAR(255) NOT NULL,
    TriggerEnabled BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT fk_device_id    FOREIGN KEY (device_id)    REFERENCES device(device_id),
    CONSTRAINT fk_sensor_id    FOREIGN KEY (sensor_id)    REFERENCES sensor(sensor_id),
    CONSTRAINT fk_attribute_id FOREIGN KEY (attribute_id) REFERENCES attribute(attribute_id),
    CONSTRAINT fk_operator_id  FOREIGN KEY (operator_id)  REFERENCES operator(id),
    CONSTRAINT fk_message_id   FOREIGN KEY (message_id)   REFERENCES notificationmessage(id)
);

CREATE TABLE IF NOT EXISTS Reminders (
    id              SERIAL PRIMARY KEY,
    message_id      INT NOT NULL,
    frequency       INT NOT NULL,
    ReminderName    VARCHAR(255) NOT NULL,
    ReminderEnabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_reminded   TIMESTAMP,
    CONSTRAINT fk_message_id FOREIGN KEY (message_id)   REFERENCES notificationmessage(id)
);

CREATE TABLE IF NOT EXISTS RemindersDevice (
    id        INT NOT NULL,
    device_id INT NOT NULL,
    rname     VARCHAR(255) NOT NULL,
    PRIMARY KEY (id, device_id),
    CONSTRAINT fk_reminders_id FOREIGN KEY (id)        REFERENCES reminders(id),
    CONSTRAINT fk_device_id    FOREIGN KEY (device_id) REFERENCES device(device_id)
);