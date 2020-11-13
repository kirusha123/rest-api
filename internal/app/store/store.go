package store

type store struct {
	cfg *sConfig
}

func New(config *sConfig) *store {
	return &store{
		cfg: config,
	}
}
