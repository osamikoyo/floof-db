package config

type Config struct {
	Addr string `toml:"address"`
}

func Get() Config {
	var cfg Config
}
