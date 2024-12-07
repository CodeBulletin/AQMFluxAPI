package config

import (
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type APIConfig struct {
	origins   string
	methods   string
	port	  string
	host 	  string
	host_html bool
	loaded    bool
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
	c.host_html = utils.GetEnv("API_HOST_HTML", "false") == "true"
	c.host = utils.GetEnv("HOST", "localhost")
	c.port = utils.GetEnv("PORT", "8080")
	c.loaded = true
}

func (c *APIConfig) Origins() string {
	return c.origins
}

func (c *APIConfig) Methods() string {
	return c.methods
}

func (c *APIConfig) Host() string {
	return c.host
}

func (c *APIConfig) Port() string {
	return c.port
}

func (c *APIConfig) HostHTML() bool {
	return c.host_html
}

func (c *APIConfig) URL() string {
	return c.Host() + ":" + c.Port()
}