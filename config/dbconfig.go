package config

import (
	"fmt"
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/utils"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	created  bool
}

func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

var dbconfig_instance *DBConfig = NewDBConfig()

func GetDBConfig() *DBConfig {
	return dbconfig_instance
}

func (c *DBConfig) Load() {

	if c.created {
		return
	}

	var port, err = strconv.Atoi(utils.GetEnv("DB_PORT", "5432"))
	if err != nil {
		port = 5432
	}

	c.Host = utils.GetEnv("DB_HOST", "localhost")
	c.Port = port
	c.User = utils.GetEnv("DB_USER", "postgres")
	c.Password = utils.GetEnv("DB_PASSWORD", "")
	c.Database = utils.GetEnv("DB_NAME", "postgres")
	c.created = true
}

func (c *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Database)
}
