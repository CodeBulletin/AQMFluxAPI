package root

import "embed"

//go:embed static
var Static embed.FS

//go:embed migrations/*.sql
var Migrations embed.FS

// GetStatic returns the static files
func GetStatic() embed.FS {
	return Static
}

// GetMigrations returns the migrations
func GetMigrations() embed.FS {
	return Migrations
}