package apiserver

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kirusha123/rest-api/internal/app/store"
	tables "github.com/kirusha123/rest-api/internal/app/store/Tables"
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

	//s.configureStore()
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
	s.router.HandleFunc("/api/create/tables", s.handleCreateTeables()).Methods("POST")
	s.router.HandleFunc("/api/create/blocks", s.handleCreateBlocks()).Methods("POST")
	s.router.HandleFunc("/api/get/blocks", s.handleGetBlocks()).Methods("GET")
	s.router.HandleFunc("/api/get/block/{id}", s.handleGetBlock()).Methods("GET")
	//s.router.HandleFunc("/api/remove/blocks", s.handleRemoveBlocks())

}

func (s *APIserver) handleGetBlock() http.HandlerFunc {
	DB := s.store.GetDB()

	var block tables.Block

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-type", "application/json")
		params := mux.Vars(r)
		n, _ := strconv.Atoi(params["id"])
		block.BlockNum = int64(n)
		err := DB.Model(&block).Where("block_num = ?block_num").Select()
		if err != nil {
			json.NewEncoder(w).Encode(block)
		} else {
			json.NewEncoder(w).Encode(tables.Block{})
			str := "params[id] = " + params["id"] + "\nQueryfailed"
			io.WriteString(w, str)
		}

	}
}

func (s *APIserver) handleGetBlocks() http.HandlerFunc {

	DB := s.store.GetDB()
	var blocks []tables.Block

	err := DB.Model(&blocks).Select()
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info(err)
			io.WriteString(w, "Query failed")
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(blocks)
		io.WriteString(w, "Query Success")
	}
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
