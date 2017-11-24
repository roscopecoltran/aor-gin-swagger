package config

import (
	"github.com/jinzhu/configor"
)

var (
	Config             ServiceConfig
	DefaultConfigFiles = []string{
		"config/database.yml",
		"config/smtp.yml",
		"config/application.yml",
		"config/cloud.yml",
		"config/plugins.yml",
		"config/tasks.yml",
	}
)

type ServiceConfig struct {
	Port  uint `env:"SERVER_PORT" default:"7000" json:"port" yaml:"port" toml:"port"`
	Debug bool `env:"SERVER_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	API   ApiConfig
	DB    DatabaseConfig
}

func init() {
	if err := configor.Load(&Config, DefaultConfigFiles...); err != nil {
		panic(err)
	}
}
