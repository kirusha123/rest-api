package tables

import (
	"math/rand"
	"time"

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
	err := DB.CreateTable(&Block{}, &orm.CreateTableOptions{IfNotExists: true})

	if err != nil {
		return err
	}

	return nil
}

//CreateTransactionTable ...
func CreateTransactionTable(DB *pg.DB) error {
	err := DB.CreateTable(&Transaction{}, &orm.CreateTableOptions{IfNotExists: true})

	if err != nil {
		return err
	}

	return nil
}

//AddBlock ...
func (b *Block) AddBlock(DB *pg.DB) error {
	err := DB.Insert(b)

	if err != nil {
		return err
	}
	return nil
}

//RemoveBlock ..
func (b *Block) RemoveBlock(DB *pg.DB) error {
	_, err := DB.Model(b).Where("block_hash = ?block_hash").Delete()
	if err != nil {
		return err
	}
	return nil
}

func CreateRandBlock(count int, DB *pg.DB) []Block {
	var blocks []Block

	rand.Seed(time.Now().Unix())

	for i := 1; i <= count; i++ {
		block := Block{
			BlockHash:        "hash",
			BlockNum:         rand.Int63(),
			TimeStamp:        rand.Int63(),
			TransactionCount: rand.Int63(),
		}
		blocks = append(blocks, block)
	}
	DB.Model(blocks).Insert()
	return blocks
}

//AddTransaction ...
func (tr *Transaction) AddTransaction(DB *pg.DB) error {
	err := DB.Insert(tr)

	if err != nil {
		return err
	}
	return nil
}
