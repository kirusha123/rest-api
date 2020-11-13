package apiserver

import "github.com/kirusha123/rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	logLevel string `toml:"log_lvl"`
	store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		logLevel: "debug",
		store:    store.NewCfg(),
	}

}
