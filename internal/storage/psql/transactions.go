package psql

import "context"

type InsertTransactionParams struct {
	Amount float64
	UserId int
	Status string
	PayId  int
	Locked bool
}

func (s *Storage) CreateTransaction(ctx context.Context, fields *InsertTransactionParams) {

}
