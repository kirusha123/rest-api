package store

type sConfig struct {
	DBURL string `toml:"db_url"`
}

func NewCfg() *sConfig {
	return &sConfig{
		DBURL: "",
	}
}
