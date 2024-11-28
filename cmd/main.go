package main

import (
	"io/fs"
	"net/http"
	"time"

	root "github.com/codebulletin/AQMFluxAPI"
	"github.com/codebulletin/AQMFluxAPI/api"
	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/mqtt"
	"github.com/codebulletin/AQMFluxAPI/ntfy"
	"github.com/codebulletin/AQMFluxAPI/server"
	"github.com/codebulletin/AQMFluxAPI/services/mqttservice"
	"github.com/codebulletin/AQMFluxAPI/services/notificationservice"
	"github.com/codebulletin/AQMFluxAPI/services/preiodic"
)


var (
	Version = "1.0.0"
	Debug = "false"
)

func main() {
	config.GetConfig().Load()

	exitChan := make(chan bool)
	freqChan := make(chan int32)

	clogger := logger.NewConsoleLogger()
	logger.SetLogger(clogger)

	logger := logger.GetLogger()

	logger.Info("Starting AQMFluxAPI Version: %s", Version)

	dbx, err := db.NewPostgresDB(logger)

	if Debug == "false" {
		fs := root.GetMigrations()
		db.MigrateUp(dbx, logger, fs)
	}

	database := db.NewDBLogger(dbx, logger)

	if err != nil {
		panic(err)
	}

	logger.DBStatus("Starting Connection to Database")

	err = db.Connect(database)

	defer db.Close(database)

	if err != nil {
		for {
			err = db.Connect(database)
			if err != nil {
				logger.Fatal("Unable to connect to database retrying in 5 secs")
				time.Sleep(5 * time.Second)
				continue
			}
			break
		}
	}

	ntfy := ntfy.New()

	mqtt_service := mqttservice.NewMqttService(database)
	notificationservice := notificationservice.NewNotificationService(ntfy)

	DBListner := db.NewPostgresListner(logger)
	defer DBListner.Close()

	DBListner.Listen("AQMFLUX_TRIGGERHIT", notificationservice.NotifyTrigger)
	DBListner.Listen("AQMFLUX_CONFIG_UPDATED", notificationservice.ChangeFreq(freqChan))

	MQTTClient := mqtt.NewMqttClient(logger)
	fetchpreodic := preiodic.NewFetchPreodic(database, MQTTClient, freqChan, exitChan)
	defer fetchpreodic.Stop()
	refreshscretes := preiodic.NewRefreshSecrets(database)
	defer refreshscretes.Stop()

	MQTTClient.Connect()
	defer MQTTClient.Disconnect()
	MQTTClient.Subscribe("esp32/input", mqtt_service.LogIntoDB)

	go DBListner.Start()
	go refreshscretes.Start()
	go fetchpreodic.Start()

	api := api.NewV1API(database)
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", api.Router()))

	if Debug == "false" {
		logger.Info("Serving Static Files")
		subFs, err := fs.Sub(root.GetStatic(), "static")
		if err != nil {
			logger.Fatal("Error getting sub filesystem: %v", err)
		}
		router.Handle("/", http.FileServer(http.FS(subFs)))
	}

	server := server.NewServer("localhost:8080", database, logger, router)
	server.Start()
	server.Close()
}
