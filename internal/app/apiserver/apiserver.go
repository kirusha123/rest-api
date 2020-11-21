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
	s.router.HandleFunc("/api/create/tables", s.handleCreateTeables())
	s.router.HandleFunc("/api/create/blocks", s.handleCreateBlocks())
	//s.router.HandleFunc("/api/remove/blocks", s.handleRemoveBlocks())
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hi, Dude")
	}
}

func (s *APIserver) handleCreateTeables() http.HandlerFunc {
	//s.store.Connect()
	err := s.store.CreateTables()
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Error to create Tables")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Tables Transaction & Blocks created")
	}

}

func (s *APIserver) handleCreateBlocks() http.HandlerFunc {
	//s.store.Connect()
	//defer s.store.Close()
	err := s.store.SetFakeBlocks()
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Blocks don't created")
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Blocks Created")
	}
}

/*func (s *APIserver) handleRemoveBlocks() http.HandlerFunc {
	//s.store.Connect()
	//defer s.store.Close()
	err := s.store.RemoveFakeBlocks()
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Blocks dont deleted")
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Blocks deleted")
	}
} //*/

func (s *APIserver) configureStore() {

	//s.store.Connect()
}
