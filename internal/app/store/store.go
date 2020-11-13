package store

type Store struct {
	cfg *Config
}

func New(config *Config) *Store {
	return &Store{
		cfg: config,
	}
}
