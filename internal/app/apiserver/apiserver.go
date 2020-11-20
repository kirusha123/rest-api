package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kirusha123/rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

//APIserver ...
type APIserver struct {
	cfg    *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//New ...
func New(config *Config) *APIserver {
	return &APIserver{
		cfg:    config,
		store:  store.New(config.store),
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ...
func (s *APIserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	s.logger.Info("starting api server")

	s.configureStore()
	if testDB := s.store.GetDB(); testDB == nil {
		s.logger.Info("Failed to connect db")
		return nil
	}

	return http.ListenAndServe(s.cfg.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	lvl, err := logrus.ParseLevel(s.cfg.logLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(lvl)

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello()) //test server function
	s.router.HandleFunc("/api/createtables", s.handleCreateTeables())

}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hi, Dude")
	}
}

func (s *APIserver) handleCreateTeables() http.HandlerFunc {
	/*err := s.store.CreateTables()
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Error to create Tables")
		}
	}*/
	s.store.Connect()
	s.store.CreateTables()
	str := s.cfg.store.Addr + " " + s.cfg.store.Pass + " " + s.cfg.store.User + " " + s.cfg.store.DBname
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, str)
	}

}

func (s *APIserver) configureStore() {

	s.store.Connect()
}
