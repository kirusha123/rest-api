package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	cfg *Config
	db  *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		cfg: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.cfg.DBURL /*"host=localhost port=5432 user=postgres password=admin dbname=TestDB sslmode=disable"*/)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
