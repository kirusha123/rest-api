package store

type Config struct {
	DBURL string `toml:"db_url"`
}

func NewCfg() *Config {
	return &Config{
		DBURL: "",
	}
}
