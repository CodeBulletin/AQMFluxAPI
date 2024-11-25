package config

import "github.com/joho/godotenv"

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

var instance *Config = NewConfig()

func GetConfig() *Config {
	return instance
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
