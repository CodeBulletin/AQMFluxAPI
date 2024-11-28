package preiodic

import (
	"context"
	"time"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/mqtt"
	"github.com/codebulletin/AQMFluxAPI/repo"
)

var fetch func()

func GetPreoidicFetch() func() {
	return fetch
}

type FetchPreodic struct {
	db 			db.DB
	logger 		logger.Logger
	mqtt 		*mqtt.Mqtt
	freqChan 	chan int32
	exitChan 	chan bool
	fetchChan 	chan bool
	duration 	int32
	durationSet bool
}

func NewFetchPreodic(db db.DB, mqtt *mqtt.Mqtt, freqChan chan int32, exitChan chan bool) *FetchPreodic {
	logger := logger.GetLogger()
	var f = FetchPreodic{
		db: db,
		logger: logger,
		mqtt: mqtt,
		freqChan: freqChan,
		exitChan: exitChan,
		fetchChan: make(chan bool),
		duration: 60,
		durationSet: false,
	}

	fetch = f.Fetch

	return &f
}

func (f *FetchPreodic) Stop() {
	f.exitChan <- true
}

func (f *FetchPreodic) fetchPreodicData() {
	defer func() {
		if r := recover(); r != nil {
			f.logger.Fatal("Recovered in FetchPreodicData: %v", r)
		}
	}()

	f.logger.Status("Fetching Preodic Data")
}

func (f *FetchPreodic) getConfigFromDB()  {
	// Get Config from DB
	query := repo.New(f.db)
	defer query.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	duration, err := query.GetIntervalConfig(ctx)

	if err != nil {
		f.logger.Error("Error while fetching config from DB: %v", err)
		return
	}

	f.duration = duration
	f.durationSet = true
}

func (f *FetchPreodic) Fetch() {
	f.fetchChan <- true
}

func (f *FetchPreodic) Start() {
	f.getConfigFromDB()
	f.fetchPreodicData()
	f.logger.Status("Preodic Fetch Started with Frequency %d", f.duration)
	for {
		select {
		case newDuration := <-f.freqChan:
			f.logger.Status("Preodic Fetch Frequency Changed to %d", newDuration)
			f.duration = newDuration
		case <-time.After(time.Duration(f.duration) * time.Second):
			f.fetchPreodicData()
			if !f.durationSet {
				f.getConfigFromDB()
			}
		case <-time.After(60 * time.Minute):
			f.getConfigFromDB()
		case <-f.fetchChan:
			f.fetchPreodicData()
		case <-f.exitChan:
			f.logger.Status("Preodic Fetch Stopped")
			return
		}
	}
}