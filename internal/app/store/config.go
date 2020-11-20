package store

//Config ...
type Config struct {
	//DBURL  string `toml:"db_url"`
	User   string `toml:"db_user"`
	Pass   string `toml:"db_pass"`
	Addr   string `toml:"db_addr"`
	DBname string `toml:"db_name"`
}

//NewCfg ...
func NewCfg() *Config {
	return &Config{
		//DBURL: "host=localhost port=5432 user=postgres password=admin dbname=TestDB sslmode=disable",
		User:   "postgres",
		Pass:   "admin",
		Addr:   "localhost:5432",
		DBname: "TestDB",
	}
}
