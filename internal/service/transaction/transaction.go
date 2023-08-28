package transaction

import (
	"context"
	"github.com/mc_transaction/internal/storage/psql"
)

type Storage interface {
	CreateTransaction(ctx context.Context, fields *psql.InsertTransactionParams)
}

type Transaction struct {
	storage Storage
}

func New(storage Storage) *Transaction {
	return &Transaction{
		storage: storage,
	}
}

type InsertParams struct {
	Amount float64
	UserId int
	Status string
	PayId  int
	Locked bool
}

func (t *Transaction) CreateTransaction(ctx context.Context, params *InsertParams) {

}
