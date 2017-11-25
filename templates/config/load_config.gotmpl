package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
)

var (
	Config             ServiceConfig
	DefaultConfigFiles = []string{
		"shared/conf.d/server.yaml",
		"shared/conf.d/api.yaml",
		"shared/conf.d/database.yaml",
		"shared/conf.d/smtp.yaml",
		"shared/conf.d/application.yaml",
		"shared/conf.d/cloud.yaml",
		"shared/conf.d/plugins.yaml",
		"shared/conf.d/tasks.yaml",
		"tests/conf.d/database.yaml",
		"tests/conf.d/smtp.yaml",
		"tests/conf.d/application.yaml",
		"tests/conf.d/cloud.yaml",
		"tests/conf.d/plugins.yaml",
		"tests/conf.d/tasks.yaml",
		"tests/conf.d/server.yaml",
		"tests/conf.d/api.yaml",
	}
)

var ConfigVar = struct {
	Port  uint           `env:"SERVER_PORT" default:"7000" json:"port" yaml:"port" toml:"port"`
	Debug bool           `env:"SERVER_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	API   ApiConfig      `json:"api" yaml:"api" toml:"api"`
	DB    DatabaseConfig `json:"db" yaml:"db" toml:"db"`
}{}

type ServiceConfig struct {
	Port  uint           `env:"SERVER_PORT" default:"7000" json:"port" yaml:"port" toml:"port"`
	Debug bool           `env:"SERVER_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	API   ApiConfig      `json:"api" yaml:"api" toml:"api"`
	DB    DatabaseConfig `json:"db" yaml:"db" toml:"db"`
	// SVC   ApiConfig      `json:"svc" yaml:"svc" toml:"svc"`
}

func init() {
	// if err := configor.New(&configor.Config{Debug: true}).Load(&Config, DefaultConfigFiles...); err != nil {
	if err := configor.Load(&Config, DefaultConfigFiles...); err != nil {
		fmt.Println("Error while loading config files: ", err)
		// panic(err)
	}
}

func PrintConfig() {
	pp.Println("Config: ", Config)
}
