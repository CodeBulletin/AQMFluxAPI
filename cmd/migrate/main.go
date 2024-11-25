package main

import (
	"log"
	"os"
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config.GetConfig().Load()

	clogger := logger.NewConsoleLogger()
	logger.SetLogger(clogger)

	database, err := db.NewPostgresDB(clogger)

	if err != nil {
		panic(err)
	}

	driver, _ := postgres.WithInstance(database, &postgres.Config{})

	m, _ := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", "postgres", driver)

	cmd := os.Args[len(os.Args)-1]

	if cmd == "down" {
		err := m.Down()

		if err != nil {
			log.Println("migration down failed", err)
			return
		}

		log.Println("Database migration - down finished")
		return
	}

	if cmd == "fix" {
		ver, err := strconv.Atoi(os.Args[len(os.Args)-2])

		if err != nil {
			log.Println("migration fix failed", err)
			return
		}

		err = m.Force(ver)

		if err != nil {
			log.Println(err)
			return
		}

		err = m.Up()

		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Database migration - fix finished")
		return
	}

	if cmd == "fixdown" {
		ver, err := strconv.Atoi(os.Args[len(os.Args)-2])

		if err != nil {
			log.Println("migration fix failed", err)
			return
		}

		err = m.Force(ver)

		if err != nil {
			log.Println(err)
			return
		}

		err = m.Down()

		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Database migration - fix finished")
		return
	}

	err = m.Up()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Database migration - up finished")
}
