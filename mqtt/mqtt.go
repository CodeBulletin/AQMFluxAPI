package mqtt

import (
	"fmt"
	"time"

	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/logger"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	opts   *MQTT.ClientOptions
	client MQTT.Client
	logger logger.Logger
}

func NewMqttClient(logger logger.Logger) *Mqtt {
	var m Mqtt = Mqtt{
		logger: logger,
	}

	var mqttconf = config.GetMQTTConfig()
	mqttconf.Load()

	var broker = fmt.Sprintf("tcp://%s:%d", mqttconf.Host, mqttconf.Port)

	m.Create(broker, "AQMFluxAPI")

	return &m
}

func (m *Mqtt) Create(broker string, client_id string) {
	m.opts = MQTT.NewClientOptions().AddBroker(broker).SetClientID(client_id)

	m.opts.SetUsername(config.GetMQTTConfig().User)
	m.opts.SetPassword(config.GetMQTTConfig().Password)

	m.opts.OnConnect = m.onConnect
	m.opts.OnConnectionLost = m.onConnectionLost

	// Set AutoReconnect to true
	m.opts.SetAutoReconnect(true)
	m.opts.SetCleanSession(true)
	m.opts.SetConnectRetry(true)
	m.opts.SetConnectRetryInterval(5 * time.Second)

	client := MQTT.NewClient(m.opts)

	m.logger.Status("MQTT Client Created")

	m.client = client
}
func (m *Mqtt) Connect() {
	m.logger.Status("Connecting to MQTT on %s", m.opts.Servers[0].String())
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		m.logger.Error("Error while connecting to MQTT %s", token.Error().Error())
	}
	m.logger.Status("Connected to MQTT")
}

func (m *Mqtt) Publish(topic string, message string) {
	token := m.client.Publish(topic, 1, false, message)
	token.Wait()
	if token.Error() != nil {
		m.logger.Error("Error while publishing %s", token.Error().Error())
	} else {
		m.logger.Info("Published to %s: %s", topic, message)
	}
}

func (m *Mqtt) Subscribe(topic string, callback func(MQTT.Client, MQTT.Message)) {
	m.logger.Status("Subscribing to %s", topic)

	m.client.Subscribe(topic, 1, callback)
}

func (m *Mqtt) onConnect(c MQTT.Client) {
	m.logger.Status("Connected to MQTT")
}

func (m *Mqtt) onConnectionLost(c MQTT.Client, err error) {
	m.logger.Error("Connection to MQTT Lost: %s %s", err.Error(), "Reconnecting")
}

func (m *Mqtt) Disconnect() {
	m.logger.Status("Disconnecting MQTT")
	m.client.Disconnect(250)
}
