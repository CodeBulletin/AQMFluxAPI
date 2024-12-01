CREATE OR REPLACE FUNCTION measurements_threshold_fn() 
RETURNS TRIGGER AS $$
DECLARE
    trigger_cursor CURSOR FOR
        SELECT o.op, o.id AS oid, t1.TriggerName AS name, t1.value1, t1.value2, t1.id FROM threshold t1
        JOIN operator o ON t1.operator_id = o.id
        WHERE t1.sensor_id = NEW.sensor_id
        AND t1.device_id = NEW.device_id
        AND t1.attribute_id = NEW.attribute_id
        AND t1.TriggerEnabled = TRUE
        AND t1.last_triggered < NOW() - INTERVAL '1 second' * t1.frequency OR t1.last_triggered IS NULL;
    trigger_row RECORD;
    has_trigger BOOLEAN;
    attr_name   TEXT;
    unit        TEXT;
    sen_name    TEXT;
    dev_name    TEXT;
    dev_loc     TEXT;
    msg_id      INT;
    time        TIMESTAMP;
BEGIN
    FOR trigger_row IN trigger_cursor LOOP
        has_trigger := FALSE;
        IF (trigger_row.op = '<') THEN
            IF NEW.mvalue < trigger_row.value1 THEN
                has_trigger := TRUE;
            END IF;
        ELSIF (trigger_row.op = '>') THEN
            IF NEW.mvalue > trigger_row.value1 THEN
                has_trigger := TRUE;
            END IF;
        ELSIF (trigger_row.op = '=') THEN
            IF NEW.mvalue = trigger_row.value1 THEN
                has_trigger := TRUE;
            END IF;
        ELSIF (trigger_row.op = '<=') THEN
            IF NEW.mvalue <= trigger_row.value1 THEN
                has_trigger := TRUE;
            END IF;
        ELSIF (trigger_row.op = '>=') THEN
            IF NEW.mvalue >= trigger_row.value1 THEN
                has_trigger := TRUE;
            END IF;
        ELSIF (trigger_row.op = '!=') THEN
            IF NEW.value != trigger_row.value1 THEN
                has_trigger := TRUE;
            END IF;
        ELSIF (trigger_row.op = 'between') THEN
            IF NEW.mvalue BETWEEN trigger_row.value1 AND trigger_row.value2 THEN
                has_trigger := TRUE;
            END IF;
        ELSIF (trigger_row.op = 'outside') THEN
            IF NEW.mvalue NOT BETWEEN trigger_row.value1 AND trigger_row.value2 THEN
                has_trigger := TRUE;
            END IF;
        END IF;

        IF has_trigger THEN
            SELECT attribute_name INTO attr_name FROM attribute WHERE attribute_id = NEW.attribute_id;
            SELECT attribute_unit INTO unit FROM attribute WHERE attribute_id = NEW.attribute_id;
            SELECT sensor_name INTO sen_name FROM sensor WHERE sensor_id = NEW.sensor_id;
            SELECT device_name INTO dev_name FROM device WHERE device_id = NEW.device_id;
            SELECT location_name INTO dev_loc FROM DeviceLocation INNER JOIN device ON DeviceLocation.location_id = device.location_id WHERE device.device_id = NEW.device_id;
            SELECT message_id INTO msg_id FROM threshold WHERE id = trigger_row.id;
            SELECT NOW() INTO time;
            PERFORM pg_notify('AQMFLUX_TRIGGERHIT', format('{"ID":%s,"0":%s,"1":%s,"2":%s,"AttrName":"%s","OP":"%s","AlertName":"%s","SenName":"%s","DevName":"%s","Time":"%s","Loc":"%s","Unit":"%s"}', msg_id, NEW.mvalue, trigger_row.value1, coalesce(trigger_row.value2, 0.0), attr_name, trigger_row.op, trigger_row.name, sen_name, dev_name, time, dev_loc, unit));
            UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
        END IF;
    END LOOP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER measurements_threshold
AFTER INSERT ON measurements
FOR EACH ROW
EXECUTE FUNCTION measurements_threshold_fn();
