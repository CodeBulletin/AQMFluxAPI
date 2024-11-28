package main

import (
	"log"
	"os"
	"strconv"

	root "github.com/codebulletin/AQMFluxAPI"
	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
)

func main() {	
	config.GetConfig().Load()

	clogger := logger.NewConsoleLogger()
	logger.SetLogger(clogger)

	logger := logger.GetLogger()

	dbx, err := db.NewPostgresDB(logger)

	if err != nil {
		log.Println(err)
		return
	}

	fs := root.GetMigrations()
	
	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		db.MigrateUp(dbx, logger, fs)
	}

	if cmd == "down" {
		db.MigrateDown(dbx, logger, fs)
	}

	if cmd == "fix-up" {
		ver := os.Args[len(os.Args)-2]
		version, err := strconv.Atoi(ver)
		if err != nil {
			log.Println(err)
			return
		}
		db.MigrateFixUp(dbx, logger, fs, version)
	}

	if cmd == "fix-down" {
		ver := os.Args[len(os.Args)-2]
		version, err := strconv.Atoi(ver)
		if err != nil {
			log.Println(err)
			return
		}
		db.MigrateFixDown(dbx, logger, fs, version)
	}
}
