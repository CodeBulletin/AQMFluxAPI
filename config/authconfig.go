package config

import (
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/utils"
)

type AUTHConfig struct {
	tokenDuration         int
	refreshTokenDuration  int
	loaded				  bool
}

func NewAUTHConfig() *AUTHConfig {
	return &AUTHConfig{}
}

var authconfig_instance *AUTHConfig = NewAUTHConfig()

func GetAUTHConfig() *AUTHConfig {
	return authconfig_instance
}

func (c *AUTHConfig) Load() {
	if c.loaded {
		return
	}
	td := utils.GetEnv("API_AUTHORIZATION_DURATION", "60")
	t1, err := strconv.Atoi(td)
	if err != nil {
		t1 = 60
	}

	rd := utils.GetEnv("API_REFRESH_DURATION", "120")
	t2, err := strconv.Atoi(rd)
	if err != nil {
		t2 = 120
	}

	c.tokenDuration = t1
	c.refreshTokenDuration = t2
	c.loaded = true
}

func (c *AUTHConfig) TokenDuration() int {
	return c.tokenDuration
}

func (c *AUTHConfig) RefreshTokenDuration() int {
	return c.refreshTokenDuration
}
