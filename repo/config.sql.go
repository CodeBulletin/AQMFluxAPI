// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: config.sql

package repo

import (
	"context"

	"github.com/lib/pq"
)

const deleteConfigByKey = `-- name: DeleteConfigByKey :one
delete from config where ckey = $1 returning ckey, cvalue, ctype
`

func (q *Queries) DeleteConfigByKey(ctx context.Context, ckey string) (Config, error) {
	row := q.queryRow(ctx, q.deleteConfigByKeyStmt, deleteConfigByKey, ckey)
	var i Config
	err := row.Scan(&i.Ckey, &i.Cvalue, &i.Ctype)
	return i, err
}

const getConfigByKey = `-- name: GetConfigByKey :one
select ckey, cvalue, ctype from config where ckey = $1
`

func (q *Queries) GetConfigByKey(ctx context.Context, ckey string) (Config, error) {
	row := q.queryRow(ctx, q.getConfigByKeyStmt, getConfigByKey, ckey)
	var i Config
	err := row.Scan(&i.Ckey, &i.Cvalue, &i.Ctype)
	return i, err
}

const getIntervalConfig = `-- name: GetIntervalConfig :one
select NULLIF(cvalue, '')::int from config where ckey = 'UPDATE INTERVAL'
`

func (q *Queries) GetIntervalConfig(ctx context.Context) (int32, error) {
	row := q.queryRow(ctx, q.getIntervalConfigStmt, getIntervalConfig)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const loadConfigData = `-- name: LoadConfigData :many
select ckey, cvalue, ctype from config where ckey = ANY($1::VARCHAR(255)[])
`

func (q *Queries) LoadConfigData(ctx context.Context, dollar_1 []string) ([]Config, error) {
	rows, err := q.query(ctx, q.loadConfigDataStmt, loadConfigData, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Config{}
	for rows.Next() {
		var i Config
		if err := rows.Scan(&i.Ckey, &i.Cvalue, &i.Ctype); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setConfigBykey = `-- name: SetConfigBykey :exec
update config 
set cvalue = data.cvalue
from (
    SELECT UNNEST($1::VARCHAR(255)[]) as ckey, UNNEST($2::TEXT[]) as cvalue
) data
where config.ckey = data.ckey
`

type SetConfigBykeyParams struct {
	Column1 []string `json:"column_1"`
	Column2 []string `json:"column_2"`
}

func (q *Queries) SetConfigBykey(ctx context.Context, arg SetConfigBykeyParams) error {
	_, err := q.exec(ctx, q.setConfigBykeyStmt, setConfigBykey, pq.Array(arg.Column1), pq.Array(arg.Column2))
	return err
}