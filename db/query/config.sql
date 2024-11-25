-- name: LoadConfigData :many
select * from config where ckey = ANY($1::VARCHAR(255)[]);

-- name: GetConfigByKey :one
select * from config where ckey = $1;

-- name: SetConfigBykey :exec
update config 
set cvalue = data.cvalue
from (
    SELECT UNNEST($1::VARCHAR(255)[]) as ckey, UNNEST($2::TEXT[]) as cvalue
) data
where config.ckey = data.ckey;

-- name: DeleteConfigByKey :one
delete from config where ckey = $1 returning *;

-- name: GetIntervalConfig :one
select NULLIF(cvalue, '')::int from config where ckey = 'UPDATE INTERVAL';