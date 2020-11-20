package tables

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

//Block ...
type Block struct {
	RefPointer int      `sql:"-"`
	TableName  struct{} `sql:"Blocks"`
	BlockNum   int64    `sql:"block_num"`
	BlockHash  string   `sql:"block_hash"`
	TimeStamp  int64    `sql:"time_stamp"`
}

//Transaction ...
type Transaction struct {
	RefPointer  int      `sql:"-"`
	TableName   struct{} `sql:"Transaction"`
	ID          int64    `sql:"id,unique,pk"`
	TXHash      string   `sql:"tx_hash,unique"`
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

//CreateBlockTable ...
func CreateBlockTable(DB *pg.DB) error {
	err := DB.CreateTable(&Block{}, &orm.CreateTableOptions{})

	if err != nil {
		return err
	}

	return nil
}

//CreateTransactionTable ...
func CreateTransactionTable(DB *pg.DB) error {
	err := DB.CreateTable(&Transaction{}, &orm.CreateTableOptions{})

	if err != nil {
		return err
	}

	return nil
}
