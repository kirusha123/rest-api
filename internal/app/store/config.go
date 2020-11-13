package store

type Config struct {
	DBURL string `toml:"db_url"`
}

func NewCfg() *Config {
	return &Config{
		DBURL: "host=localhost port=5432 user=postgres password=admin dbname=TestDB sslmode=disable",
	}
}
