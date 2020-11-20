package store

import (
	"github.com/go-pg/pg"
	tables "github.com/kirusha123/rest-api/internal/app/store/Tables"
	//_ "github.com/lib/pq"
)

//Store ...
type Store struct {
	cfg *Config
	db  *pg.DB
}

//New ...
func New(config *Config) *Store {
	return &Store{
		cfg: config,
		db:  nil,
	}
}

//GetDB ...
func (s *Store) GetDB() *pg.DB {
	return s.db
}

//Connect ...
func (s *Store) Connect() {
	/*db, err := sql.Open("postgres", s.cfg.DBURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db*/

	s.db = pg.Connect(&pg.Options{
		User:     s.cfg.User,
		Password: s.cfg.Pass,
		Addr:     s.cfg.Addr,
		Database: s.cfg.DBname,
	})

}

//Close ...
func (s *Store) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}

//CreateTables ...
func (s *Store) CreateTables() error {
	err := tables.CreateBlockTable(s.db)
	if err != nil {
		return err
	}
	err = tables.CreateTransactionTable(s.db)
	if err != nil {
		return err
	}

	return nil
}
