package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kirusha123/rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	cfg    *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		cfg:    config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store.New(config.store),
	}
}

func (s *APIserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	s.logger.Info("starting api server")

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
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hi, Dude")
	}
}
