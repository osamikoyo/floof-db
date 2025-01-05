package config

import (
	"github.com/BurntSushi/toml"
	"github.com/osamikoyo/floof-db/pkg/loger"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type Config struct {
	Addr string `toml:"address"`
}

func Get() Config {
	var cfg Config

	home, _ := os.UserHomeDir()
	configs, _ := os.UserConfigDir()

	if _, err := toml.DecodeFile(filepath.Join(home, ".floof", "config.toml"), &cfg); err != nil {
		if _, err = toml.DecodeFile(filepath.Join(configs, "floof", "config.toml"), &cfg); err != nil {
			loger.New().Error().Err(errors.New("Could not find a config files in Dirs"))
			return cfg
		}
	}

	return cfg
}
