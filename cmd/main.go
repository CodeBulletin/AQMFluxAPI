package main

import (
	"time"

	"github.com/codebulletin/AQMFluxAPI/cmd/api"
	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/mqtt"
	"github.com/codebulletin/AQMFluxAPI/ntfy"
	"github.com/codebulletin/AQMFluxAPI/services/mqttservice"
	"github.com/codebulletin/AQMFluxAPI/services/notificationservice"
	"github.com/codebulletin/AQMFluxAPI/services/preiodic"
)

func main() {
	config.GetConfig().Load()

	exitChan := make(chan bool)
	freqChan := make(chan int32)

	clogger := logger.NewConsoleLogger()
	logger.SetLogger(clogger)

	logger := logger.GetLogger()

	dbx, err := db.NewPostgresDB(logger)

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

	server := api.NewServer("localhost:8080", database, logger)

	server.Start()
}
