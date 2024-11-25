package config

import (
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type NTFYConfig struct {
	Host  string
	Token string
}

func NewNTFYConfig() *NTFYConfig {
	return &NTFYConfig{}
}

var ntfyconfig_instance *NTFYConfig = NewNTFYConfig()

func GetNTFYConfig() *NTFYConfig {
	return ntfyconfig_instance
}

func (c *NTFYConfig) Load() {
	c.Host = utils.GetEnv("NTFY_HOST", "localhost")
	c.Token = utils.GetEnv("NTFY_TOKEN", "")
}
