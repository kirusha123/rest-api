package apiserver

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kirusha123/http-rest-api/iternal/app/store"
	tables "github.com/kirusha123/http-rest-api/iternal/app/store/table"
	"github.com/sirupsen/logrus"
)

//APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//Start ....
func (s *APIserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	dt := time.Now()
	s.logger.Info("starting apiserver\nServer Time: ", dt.String())

	return http.ListenAndServe(s.config.bindAddr, s.router)
}

//New ...
func New(conf *Config) *APIserver {
	return &APIserver{
		config: conf,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store.New(conf.store),
	}
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.logLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter() {

	s.router.HandleFunc("/hello", s.handlehello())
	s.router.HandleFunc("/api/get/time", s.handletime()).Methods("GET")
	s.router.HandleFunc("/api/get/txs", s.handleGetTxs()).Methods("GET")
	s.router.HandleFunc("/api/get/blocks", s.handleGetBlocks()).Methods("GET")
	s.router.HandleFunc("/api/get/blocks/{id}", s.handleGetBlock()).Methods("GET")
	s.router.HandleFunc("/api/get/blocktxs/{bid}", s.handleGetBlockTxs()).Methods("GET")
	s.router.HandleFunc("/api/create/tables", s.handlecCrtTbls()).Methods("POST")
	s.router.HandleFunc("/api/create/blocks/{count}", s.handleCreateBlocksN()).Methods("POST")
	s.router.HandleFunc("/api/create/transactions/{Blockid}", s.handleCreateTransactions()).Methods("POST")
	s.router.HandleFunc("/api/create/transaction/{bid}", s.handleCreateTx()).Methods("POST")
}

func (s *APIserver) handleGetTxs() http.HandlerFunc {
	DB := s.store.GetDB()
	if DB == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info("Failed connect to db")
			io.WriteString(w, "Failed connect to db")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var txs []*tables.Transaction
		q_err := DB.Model(&txs).Select()
		if q_err != nil {
			s.logger.Info("\nFailed to get TXs\n", q_err)
		} else {
			json.NewEncoder(w).Encode(txs)
		}
	}
}

func (s *APIserver) handleCreateTx() http.HandlerFunc {
	DB := s.store.GetDB()
	if DB == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info("Failed connect to db")
			io.WriteString(w, "Failed connect to db")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		id, err := strconv.ParseInt(mux.Vars(r)["bid"], 10, 0)
		if err != nil {
			s.logger.Info("\nfailed to get bid \n", err)
		}

		tx := &tables.Transaction{
			TXHash:      "TxHash",
			ID:          rand.Int63(),
			BlockNum:    int64(id),
			BlockHash:   "BlockHash",
			TimeStamp:   time.Now().Unix(),
			IsCoinbase:  true,
			InputCount:  rand.Int63n(100000),
			InputValue:  rand.Int63n(10000),
			OutputCount: rand.Int63n(100000),
			OutputValue: rand.Int63n(10000),
			Fee:         rand.Int63n(100),
		}

		_, qin_err := DB.Model(tx).Insert()
		if qin_err != nil {
			s.logger.Info("\nfailed to add TX\n", qin_err)
		} else {
			b := &tables.Block{
				BlockNum: id,
			}
			qs_err := DB.Model(b).Where("block_num = ?block_num").Select()
			if qs_err != nil {
				s.logger.Info("\nfailed to Select Block with block_id = ", id, "\n", qs_err)
			} else {
				b_ch := &tables.Block{
					BlockNum:         b.BlockNum,
					TransactionCount: int64(b.TransactionCount + 1),
				}
				_, qup_err := DB.Model(b_ch).Set("transaction_count = ?transaction_count").Where("block_num = ?block_num").Update()
				if qup_err != nil {
					s.logger.Info("Failed to Update Block with block_id = ", id, "\n", qup_err)
				} else {
					s.logger.Info("TX added successfully")
				}
			}

		}

	}
}

func (s *APIserver) handleGetBlockTxs() http.HandlerFunc {

	DB := s.store.GetDB()
	if DB == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info("Failed connect to db")
			io.WriteString(w, "Failed connect to db")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		id, err := strconv.ParseInt(mux.Vars(r)["bid"], 10, 0)
		if err != nil {
			s.logger.Info("\nfailed to get bid \n", err)
		}
		var txs []*tables.Transaction

		q_err := DB.Model(&txs).Where("block_num = ?", id).Select()

		if q_err != nil {
			s.logger.Info("\nFaied to get txs\n", err)
		} else {
			json.NewEncoder(w).Encode(txs)
		}

	}
}

func (s *APIserver) handleGetBlock() http.HandlerFunc {
	DB := s.store.GetDB()
	if DB == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info("Failed connect to db")
			io.WriteString(w, "Failed connect to db")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.logger.Info("\nfailed to get id \n", err)
		}

		block := &tables.Block{
			BlockNum: int64(id),
		}

		sel_err := DB.Model(block).Where("block_num = ?block_num").Select()
		if sel_err != nil {
			s.logger.Info("failed to get Block from db\n", sel_err)
		} else {
			json.NewEncoder(w).Encode(block)
		}

	}
}

func (s *APIserver) handleGetBlocks() http.HandlerFunc {
	DB := s.store.GetDB()
	if DB == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info("Failed connect to db")
			io.WriteString(w, "Failed connect to db")
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var blocks []*tables.Block
		err := DB.Model(&blocks).Select()
		if err != nil {
			s.logger.Info("Request errr:\n", err)
		}
		w.Header().Set("Content-type", "application/json")
		if len(blocks) != 0 {
			json.NewEncoder(w).Encode(blocks)
		} else {
			json.NewEncoder(w).Encode(&tables.Block{})
		}
	}
}

func (s *APIserver) handleCreateTransactions() http.HandlerFunc {
	DB := s.store.GetDB()
	if DB == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Failed connect to db")
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		params := mux.Vars(r)

		blockId, _ := strconv.Atoi(params["Blockid"])
		block := &tables.Block{
			BlockNum: int64(blockId),
		}
		check, err := DB.Model(block).Where("block_num = ?block_num").Exists()
		if err != nil {
			s.logger.Info("Failed to check block\n", err)
		} else {
			if check == true {
				err := DB.Model(block).Where("block_num = ?block_num").Select()
				if err != nil {
					s.logger.Info("Selectiont of Block failed \n", err)
				}
				s.logger.Info(block)
				txCount := block.TransactionCount
				var i int64
				rand.Seed(time.Now().Unix())
				s.logger.Info("I'm here\n\ntx_count = ", txCount)
				for i = 1; i <= txCount; i++ {
					tx := &tables.Transaction{
						TXHash:      "TxHash",
						ID:          rand.Int63(),
						BlockNum:    int64(blockId),
						BlockHash:   "BlockHash",
						TimeStamp:   time.Now().Unix(),
						IsCoinbase:  true,
						InputCount:  rand.Int63n(100000),
						InputValue:  rand.Int63n(10000),
						OutputCount: rand.Int63n(100000),
						OutputValue: rand.Int63n(10000),
						Fee:         rand.Int63n(100),
					}
					_, err := DB.Model(tx).Insert()
					if err != nil {
						s.logger.Info(err, "\nerr_num = ", i)
						io.WriteString(w, "Querry failed")
						break
					}
				}

			} else {
				s.logger.Info("Block doesn't exists\n\nblock_num = ", blockId)
			}

		}
	}
}

func (s *APIserver) handleCreateBlocksN() http.HandlerFunc {

	DB := s.store.GetDB()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		params := mux.Vars(r)
		count, _ := strconv.Atoi(params["count"])
		rand.Seed(time.Now().Unix())
		for i := 1; i <= count; i++ {
			b := &tables.Block{
				BlockHash:        "hash",
				BlockNum:         rand.Int63n(650000),
				TimeStamp:        rand.Int63(),
				TransactionCount: rand.Int63n(10),
			}
			_, err := DB.Model(b).Insert()
			if err != nil {
				s.logger.Info(err, "\nerr_num = ", i)
				io.WriteString(w, "Querry failed")
				break
			}
		}

		/*blocks = tables.CreateRandBlock(count, DB)
		if len(blocks) != 0 {
			json.NewEncoder(w).Encode(blocks)
		} else {
			json.NewEncoder(w).Encode(tables.Block{})
		}*/
		io.WriteString(w, "function ended")
	}
}

func (s *APIserver) handlecCrtTbls() http.HandlerFunc {
	DB := s.store.GetDB()

	if DB == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Failed to create db")
		}
	}

	if err := tables.CreateBlockTable(DB); err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info(err)
			io.WriteString(w, "Failed to create Block db")
		}
	}

	if err := tables.CreateTxTable(DB); err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			s.logger.Info(err)
			io.WriteString(w, "Failed to create Transaction db")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "DB created")
	}
}

func (s *APIserver) handletime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		dt := time.Now()
		json.NewEncoder(w).Encode(dt)
	}
}

func (s *APIserver) handlehello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Dude")
	}
}
