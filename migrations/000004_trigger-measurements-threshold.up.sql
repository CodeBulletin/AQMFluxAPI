CREATE OR REPLACE FUNCTION measurements_threshold_fn() 
RETURNS TRIGGER AS $$
DECLARE
    trigger_cursor CURSOR FOR
        SELECT o.op, t1.value1, t1.value2, t1.id FROM threshold t1
        JOIN operator o ON t1.operator_id = o.id
        WHERE t1.sensor_id = NEW.sensor_id
        AND t1.device_id = NEW.device_id
        AND t1.attribute_id = NEW.attribute_id
        AND t1.TriggerEnabled = TRUE
        AND t1.last_triggered < NOW() - INTERVAL '1 second' * t1.frequency OR t1.last_triggered IS NULL;
    trigger_row RECORD;
BEGIN
    FOR trigger_row IN trigger_cursor LOOP
        IF (trigger_row.op = '<') THEN
            IF NEW.value < trigger_row.value1 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        ELSIF (trigger_row.op = '>') THEN
            IF NEW.value > trigger_row.value1 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        ELSIF (trigger_row.op = '=') THEN
            IF NEW.value = trigger_row.value1 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        ELSIF (trigger_row.op = '<=') THEN
            IF NEW.value <= trigger_row.value1 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        ELSIF (trigger_row.op = '>=') THEN
            IF NEW.value >= trigger_row.value1 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        ELSIF (trigger_row.op = '!=') THEN
            IF NEW.value != trigger_row.value1 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        ELSIF (trigger_row.op = 'between') THEN
            IF NEW.value BETWEEN trigger_row.value1 AND trigger_row.value2 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        ELSIF (trigger_row.op = 'outside') THEN
            IF NEW.value NOT BETWEEN trigger_row.value1 AND trigger_row.value2 THEN
                PERFORM pg_notify('AQMFLUX_TRIGGERHIT', 'ID: ' || trigger_row.id);

                UPDATE threshold SET last_triggered = NOW() WHERE id = trigger_row.id;
            END IF;
        END IF;
    END LOOP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER measurements_threshold
AFTER INSERT ON measurements
FOR EACH ROW
EXECUTE FUNCTION measurements_threshold_fn();
