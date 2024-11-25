CREATE OR REPLACE FUNCTION config_updated_fn()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the ckey is 'UPDATE INTERVAL'
    IF (NEW.ckey = 'UPDATE INTERVAL') THEN
        -- Notify the user that the update interval has been changed
        PERFORM pg_notify('AQMFLUX_CONFIG_UPDATED', NEW.cvalue);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER config_updated
AFTER INSERT OR UPDATE OR DELETE ON config
FOR EACH ROW
EXECUTE FUNCTION config_updated_fn();