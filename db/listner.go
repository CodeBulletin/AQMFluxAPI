package db

import (
	"time"

	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/logger"
	pq "github.com/lib/pq"
)

type PostgresListener struct {
	Callbacks map[string]func(data string)
	Listener  *pq.Listener
	logger    logger.Logger
}

func (p *PostgresListener) eventCallback(ev pq.ListenerEventType, err error) {
	if err != nil {
		p.logger.DBError("Received an error on listener: %s", err)
	}

	switch ev {
		case pq.ListenerEventConnected:
			p.logger.DBStatus("Listener connected to PostgreSQL.")
		case pq.ListenerEventDisconnected:
			p.logger.DBError("Listener disconnected. Attempting to reconnect...")
		case pq.ListenerEventReconnected:
			p.logger.DBStatus("Listener successfully reconnected.")
	}
}

func NewPostgresListner(logger logger.Logger) *PostgresListener {
	var dbconfig = config.GetDBConfig()
	dbconfig.Load()

	var conninfo = dbconfig.GetConnectionString()
	
	logger.DBStatus("Listener created")
	
	pglistner := &PostgresListener{
		Callbacks: make(map[string]func(data string)),
		logger:    logger,
	}

	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, pglistner.eventCallback)
	pglistner.Listener = listener
	
	return pglistner
}

func (p *PostgresListener) Listen(channel string, callback func(data string)) {
	p.logger.DBStatus("Listening to channel %s", channel)
	p.Listener.Listen(channel)
	p.Callbacks[channel] = callback
}

func (p *PostgresListener) Unlisten(channel string) {
	p.Listener.Unlisten(channel)
	delete(p.Callbacks, channel)
}


func (p *PostgresListener) Start() {
	p.logger.DBStatus("Listening for notifications")
	for {
		select {
		case notification := <-p.Listener.Notify:
			if (notification == nil) {
				continue
			}
			p.logger.DBInfo("Received data from channel %s : %s\n", notification.Channel, notification.Extra)
			p.Callbacks[notification.Channel](notification.Extra)
		case <-time.After(5 * time.Minute):
			p.logger.DBStatus("No notifications for 5 minutes. Sending ping...")
			go func() {
				err := p.Listener.Ping()
				if err != nil {
					p.logger.Error("Ping error: %v", err)
					return
				}
				p.logger.DBStatus("Ping Successfull")
			}()
		}
	}
}

func (p *PostgresListener) Close() {
	p.logger.DBStatus("Closing listener")
	err := p.Listener.Close()
	if err != nil {
		p.logger.DBError("Error closing listener: %s\n", err)
	}
}

func (p *PostgresListener) Ping() {
	p.logger.DBStatus("Pinging listener")
	err := p.Listener.Ping()
	if err != nil {
		p.logger.DBError("Error pinging listener: %s\n", err)
	}
}