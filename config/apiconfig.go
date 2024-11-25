package config

import (
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type APIConfig struct {
	origins string
	methods string
	loaded  bool
}

func NewAPIConfig() *APIConfig {
	return &APIConfig{}
}

var apiconfig_instance *APIConfig = NewAPIConfig()

func GetAPIConfig() *APIConfig {
	return apiconfig_instance
}

func (c *APIConfig) Load() {
	if c.loaded {
		return
	}
	c.origins = utils.GetEnv("API_ORIGINS", "*")
	c.methods = utils.GetEnv("API_METHODS", "GET POST")
	c.loaded = true
}

func (c *APIConfig) Origins() string {
	return c.origins
}

func (c *APIConfig) Methods() string {
	return c.methods
}
