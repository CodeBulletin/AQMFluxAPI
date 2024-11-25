package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/codebulletin/AQMFluxAPI/logger"
)

type DBLogger struct{
	logger logger.Logger
	*sql.DB;
}

func NewDBLogger(db *sql.DB, logger logger.Logger) DBLogger {
	return DBLogger{DB: db, logger: logger}
}

func (d DBLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	now := time.Now()
	row, err := d.DB.Query(query, args...)
	if err != nil {
		d.logger.DBError("exec: %s | Query: %s", time.Since(now), query)
	} else {
		d.logger.DBInfo("exec: %s | Query: %s", time.Since(now), query)
	}
	return row, err
}

func (d DBLogger) QueryRow(query string, args ...interface{}) *sql.Row {
	now := time.Now()
	row := d.DB.QueryRow(query, args...)
	d.logger.Info("exec: %s | QueryRow: %s", time.Since(now), query)
	return row
}

func (d DBLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	now := time.Now()
	result, err := d.DB.Exec(query, args...)
	if err != nil {
		d.logger.DBError("exec: %s | Exec: %s", time.Since(now), query)
	} else {
		d.logger.DBInfo("exec: %s | Exec: %s", time.Since(now), query)
	}
	return result, err
}

func (d DBLogger) Prepare(query string) (*sql.Stmt, error) {
	now := time.Now()
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		d.logger.DBError("exec: %s | Prepare: %s", time.Since(now), query)
	} else {
		d.logger.DBInfo("exec: %s | Prepare: %s", time.Since(now), query)
	}
	return stmt, err
}

func (d DBLogger) Begin() (*sql.Tx, error) {
	now := time.Now()
	tx, err := d.DB.Begin()
	if err != nil {
		d.logger.DBFatal("exec: %s | Begin: %s", time.Since(now), err.Error())
	} else {
		d.logger.DBStatus("exec: %s | Begin: %+v", time.Since(now), tx)
	}
	return tx, err
}

func (d DBLogger) Ping() error {
	now := time.Now()
	err := d.DB.Ping()
	if err != nil {
		d.logger.DBFatal("exec: %s | Ping: %s", time.Since(now), err.Error())
	} else {
		d.logger.DBStatus("exec: %s | Ping: %s", time.Since(now), "Pong")
	}
	return err
}

func (d DBLogger) Close() error {
	now := time.Now()
	err := d.DB.Close()
	if err != nil {
		d.logger.DBFatal("exec: %s | Close: %s", time.Since(now), err.Error())
	} else {
		d.logger.DBStatus("exec: %s | Close: %s", time.Since(now), "Closed")
	}
	return err
}

func (d DBLogger) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	now := time.Now()
	row, err := d.DB.QueryContext(ctx, query, args...)
	if err != nil {
		d.logger.DBError("exec: %s | QueryContext: %sError: %s", time.Since(now), query, err.Error())
	} else {
		d.logger.DBInfo("exec: %s | QueryContext: %s", time.Since(now), query)
	}
	return row, err
}

func (d DBLogger) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	now := time.Now()
	row := d.DB.QueryRowContext(ctx, query, args...)
	d.logger.DBInfo("exec: %s | QueryRowContext: %s", time.Since(now), query)
	return row
}

func (d DBLogger) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	now := time.Now()
	result, err := d.DB.ExecContext(ctx, query, args...)
	if err != nil {
		d.logger.DBError("exec: %s | ExecContext: %sError: %s", time.Since(now), query, err.Error())
	} else {
		d.logger.DBInfo("exec: %s | ExecContext: %s", time.Since(now), query)
	}
	return result, err
}

func (d DBLogger) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	now := time.Now()
	stmt, err := d.DB.PrepareContext(ctx, query)
	if err != nil {
		d.logger.DBError("exec: %s | PrepareContext: %sError: %s", time.Since(now), query, err.Error())
	} else {
		d.logger.DBInfo("exec: %s | PrepareContext: %s", time.Since(now), query)
	}
	return stmt, err
}