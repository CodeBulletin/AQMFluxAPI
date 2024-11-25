package db

import (
	"context"
	"database/sql"

	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/logger"
	_ "github.com/lib/pq"
)

type DB interface {
	Close() error
	Ping() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
	Begin() (*sql.Tx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}


func NewPostgresDB(logger logger.Logger) (*sql.DB, error) {
	var dbconfig = config.GetDBConfig()
	dbconfig.Load()

	var url = dbconfig.GetConnectionString()

	db, err := sql.Open("postgres", url)

	if err != nil {
		logger.Error("Error opening database connection: %v", err)

		return nil, err
	}

	return db, nil
}

func Connect(db DB) error {
	return db.Ping()
}

func Close(db DB) error {
	return db.Close()
}
