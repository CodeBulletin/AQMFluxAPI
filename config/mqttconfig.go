package config

import (
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/utils"
)

type MQTTConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func NewMQTTConfig() *MQTTConfig {
	return &MQTTConfig{}
}

var mqttconfig_instance *MQTTConfig = NewMQTTConfig()

func GetMQTTConfig() *MQTTConfig {
	return mqttconfig_instance
}

func (c *MQTTConfig) Load() {
	var port, err = strconv.Atoi(utils.GetEnv("MQTT_PORT", "5432"))
	if err != nil {
		port = 5432
	}

	c.Host = utils.GetEnv("MQTT_HOST", "localhost")
	c.Port = port
	c.User = utils.GetEnv("MQTT_USER", "mqtt")
	c.Password = utils.GetEnv("MQTT_PASSWORD", "")
}
