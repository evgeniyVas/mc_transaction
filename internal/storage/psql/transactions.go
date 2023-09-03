package psql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type TransactionStorage struct {
	conn *sqlx.DB
}

func NewTransactionStorage(conn *sqlx.DB) *TransactionStorage {
	return &TransactionStorage{conn: conn}
}

type InsertTransactionParams struct {
	Amount    float64
	UserId    int64
	Status    string
	PayId     int64
	Locked    bool
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *TransactionStorage) CreateTransaction(ctx context.Context, fields *InsertTransactionParams) (int64, error) {

}
