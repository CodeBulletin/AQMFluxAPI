package mqttservice

import (
	"bytes"
	"context"
	"log"
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/db/repo"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/services/preiodic"
	"github.com/codebulletin/AQMFluxAPI/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MqttService is a struct that holds the MQTT client options and the MQTT client
type MqttService struct {
	db db.DB
	logger logger.Logger
}

// NewMqttService creates a new MqttService
func NewMqttService(db db.DB) *MqttService {
	logger := logger.GetLogger()
	return &MqttService{
		db: db,
		logger: logger,
	}
}

// Log the message
func (m *MqttService) LogIntoDB(MQTT mqtt.Client, message mqtt.Message) {
	defer func ()  {
		if r := recover(); r != nil {
			m.logger.Fatal("Recovered in LogIntoDB: %v", r)
		}
	}()

	msg := message.Payload()
	// log.Printf("TOPIC: %s : MESSAGE: %s\n", message.Topic(), string(msg))
	m.logger.Info("TOPIC: %s : MESSAGE: %s\n", message.Topic(), string(msg))

	if len(msg) < 3 {
		m.logger.Error("Invalid message: %s", string(msg))
		return
	}

	// message is of the types r:data, m:data where r is the readings and m is the message
	msgType := string(msg[0:2])
	msgData := msg[2:]

	switch msgType {
	case "r:":
		m.reading(msgData)
	case "m:":
		m.message(msgData)
	case "c:":
		m.command(msgData)
	default:
		m.logger.Error("Unknown message type: %s", msgType)
	}
}

func (m *MqttService) reading(msgData []byte) {

	if len(msgData) < 20 {
		m.logger.Error("Invalid reading message: %s", string(msgData))
		return
	}

	time := string(msgData[0:19])
	
	deviceID, nextIdx, err := utils.ExtractFirstNumber(msgData[20:])
	if err != nil {
		m.logger.Error("Error parsing device id: %v", err)
		return
	}

	if len(msgData) < nextIdx+20+2 {
		m.logger.Error("Invalid reading message: %s", string(msgData))
		return
	}

	data := msgData[nextIdx+20+2:]

	// data is of type key:value/key:value/... where key is a integer and value array of bytes

	ctx := context.Background()
	Query := repo.New(m.db)
	defer Query.Close()
	for _, part := range bytes.Split(data, []byte("/")) {

		SensorID, nextIdx, err := utils.ExtractFirstNumber(part)

		if err != nil {
			m.logger.Error("Error parsing sensor id: %v", err)
			continue
		}

		if len(part) < nextIdx+2 {
			m.logger.Error("Invalid sensor value: %s", string(part))
			continue
		}

		value := part[nextIdx+2:]

		// value is of type attr=val,attr=val,... where value is a float and attr is a string

		for _, v := range bytes.Split(value, []byte(",")) {
			attrs := bytes.Split(v, []byte("="))

			if len(attrs) != 2 {
				m.logger.Error("Invalid attribute value: %s", string(v))
				continue
			}

			attr, val := attrs[0], attrs[1]

			attrStr := string(attr)
			valStr := string(val)

			valFloat, err := strconv.ParseFloat(valStr, 64)

			if err != nil {
				m.logger.Error("Error parsing value: %v", err)
				continue
			}

			// Convert time to time.Time
			time, err := utils.ParseTime(time)

			if err != nil {
				m.logger.Error("Error parsing time: %v", err)
				continue
			}

			id, err := Query.AttributeIdFromName(ctx, attrStr)

			if err != nil {
				m.logger.Error("Error getting attribute id: %v for attribute: %s", err, attrStr)
				continue
			}

			tx, err := m.db.Begin()
			if err != nil {
				m.logger.Error("Error starting transaction: %v", err)
				continue
			}


			Qtx := Query.WithTx(tx)

			err = Qtx.InsertMeasurement(ctx, repo.InsertMeasurementParams{
				Mtime:       time,
				DeviceID:    deviceID,
				SensorID:    SensorID,
				AttributeID: id,
				Mvalue:      valFloat,
			})

			if err != nil {
				m.logger.Error("Error inserting measurement: %v", err)
				continue
			}

			err = tx.Commit()

			if err != nil {
				m.logger.Error("Error committing transaction: %v", err)
			}
		}
	}

}

func (m *MqttService) message(msgData []byte) {
	if len(msgData) < 20 {
		m.logger.Error("Invalid reading message: %s", string(msgData))
		return
	}

	time := string(msgData[0:19])
	
	deviceID, nextIdx, err := utils.ExtractFirstNumber(msgData[20:])
	if err != nil {
		m.logger.Error("Error parsing device id: %v", err)
		return
	}

	if len(msgData) < nextIdx+20+2 {
		m.logger.Error("Invalid reading message: %s", string(msgData))
		return
	}

	data := msgData[nextIdx+20+2:]

	for _, part := range bytes.Split(data, []byte("/")) {
		log.Printf("Message: TIME: %s DEVICE: %d MESSAGE: %s\n", time, deviceID, string(part))
	}
}

func (m *MqttService) command(msgData []byte) {
	defer func ()  {
		if r := recover(); r != nil {
			m.logger.Fatal("Recovered in command: %v", r)
		}
	}()

	msg := string(msgData)

	switch msg {
		case "fetch":
			fetch := preiodic.GetPreoidicFetch()
			fetch()
		default:
			m.logger.Error("Unknown command: %s", msg)
	}
}