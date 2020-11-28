package tables

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

//Block ...
type Block struct {
	RefPointer       int      `sql:"-" json:"-"`
	TableName        struct{} `sql:"Blocks" json:"-"`
	BlockNum         int64    `sql:"block_num"`
	BlockHash        string   `sql:"block_hash"`
	TimeStamp        int64    `sql:"time_stamp"`
	TransactionCount int64    `sql:"transaction_count"`
}

//Transaction ...
type Transaction struct {
	RefPointer  int      `sql:"-" json:"-"`
	TableName   struct{} `sql:"Transaction" json:"-"`
	ID          int64    `sql:"id"`
	TXHash      string   `sql:"tx_hash"`
	BlockNum    int64    `sql:"block_num"`
	BlockHash   string   `sql:"block_hash"`
	TimeStamp   int64    `sql:"time_stamp"`
	IsCoinbase  bool     `sql:"is_coinbase"`
	InputCount  int64    `sql:"input_count"`
	OutputCount int64    `sql:"output_count"`
	InputValue  int64    `sql:"input_value"`
	OutputValue int64    `sql:"output_value"`
	Fee         int64    `sql:"fee"`
}

func CreateBlockTable(DB *pg.DB) error {
	err := DB.CreateTable(&Block{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	return nil
}

func CreateTxTable(DB *pg.DB) error {
	err := DB.CreateTable(&Transaction{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	return nil
}
