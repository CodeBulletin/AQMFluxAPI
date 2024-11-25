BEGIN;

INSERT INTO config (ckey, cvalue, ctype) VALUES 
('LOCATION', 'New York', 1),
('LAT', '40.7128', 2),
('LON', '-74.0060', 2),
('OPEN WEATHER MAP API KEY', '', 0),
('OPEN METEO API KEY', '', 0),
('WEATHER UPDATE INTERVAL', '', 4),
('UPDATE INTERVAL', '15', 3),
('USERNAME', 'admin', 0),
('PASSWORD', 'admin', 0)
ON CONFLICT DO NOTHING;

COMMIT;

BEGIN;

INSERT INTO operator (id, op, variables) VALUES 
    (1, '<', 1),
    (2, '<=', 1),
    (3, '>', 1),
    (4, '>=', 1),
    (5, '=', 1),
    (6, '!=', 1),
    (7, 'between', 2),
    (8, 'outside', 2)
ON CONFLICT DO NOTHING;

COMMIT;