package config

import (
	"github.com/BurntSushi/toml"
	"github.com/osamikoyo/floof-db/pkg/loger"
)

type Config struct {
	Addr string `toml:"address"`
}

func Get() Config {
	var cfg Config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		loger.New().Error().Err(err)
	}

	return cfg
}
