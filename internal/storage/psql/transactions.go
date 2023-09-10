package psql

import (
	"context"
	"fmt"
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

func (t *TransactionStorage) CreateTransaction(ctx context.Context, p *InsertTransactionParams) (int64, error) {
	var id int64

	err := t.conn.GetContext(ctx, &id, queryInsertTransaction, p.UserId, p.Amount, p.Status, p.PayId, p.Locked, p.Token)
	if err != nil {
		return 0, fmt.Errorf("CreateTransaction.queryInsertTransaction error: %w", err)
	}

	return id, nil
}

func (t *TransactionStorage) SelectTransactionWithLock(ctx context.Context) (int64, error) {
	return 1, nil
}

func (t *TransactionStorage) UpdateTransactionStatus(ctx context.Context) error {
	return nil
}
