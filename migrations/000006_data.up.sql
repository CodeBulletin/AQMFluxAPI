BEGIN;

INSERT INTO secrets (name, value, expires_at, created_at, updated_at) VALUES 
('jwt_secret', 'secret', NOW() - INTERVAL '1 day', NOW(), NOW()),
('jwt_refresh_secret', 'refresh_secret', NOW() - INTERVAL '1 day', NOW(), NOW());

COMMIT;