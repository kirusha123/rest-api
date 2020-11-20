package apiserver

import "github.com/kirusha123/rest-api/internal/app/store"

//Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	logLevel string `toml:"log_lvl"`
	store    *store.Config
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		logLevel: "debug",
		store:    store.NewCfg(),
	}

}
