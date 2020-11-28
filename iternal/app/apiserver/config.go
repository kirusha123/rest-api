package apiserver

import "github.com/kirusha123/http-rest-api/iternal/app/store"

//Config ...
type Config struct {
	bindAddr string `toml:"bind_addr"`
	logLevel string `toml:"log_level"`
	store    *store.Config
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		bindAddr: ":8080",
		logLevel: "debug",
		store:    store.NewCfg(),
	}
}
