package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (t *TransactionStorage) CreateTransaction(ctx context.Context, p *InsertTransactionParams) error {
	_, err := t.conn.ExecContext(ctx, queryInsertTransaction, p.UserId, p.Amount, p.Status, p.PayId, p.Locked, p.Token)
	if err != nil {
		return fmt.Errorf("CreateTransaction.queryInsertTransaction error: %w", err)
	}

	return nil
}

var ErrTransactionNotFound = errors.New("transaction not found")

type Transaction struct {
	ID        int64     `db:"id"`
	PayId     int64     `db:"pay_id"`
	Amount    float64   `db:"amount"`
	UserId    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (t *TransactionStorage) SelectTransactionWithLock(ctx context.Context) (*Transaction, error) {
	var (
		tx  *sqlx.Tx
		err error
	)
	defer rollBackTx(tx, err)

	tx, err = t.conn.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("tx begin error: %w", err)
	}

	res := &Transaction{}
	err = tx.GetContext(ctx, res, querySelectTransactionWithLock)
	if errors.Is(err, sql.ErrNoRows) {
		err = tx.Commit()
		return nil, ErrTransactionNotFound
	} else if err != nil {
		return nil, fmt.Errorf("select error %w", err)
	}

	_, err = tx.ExecContext(ctx, queryUpdateTransactionLocked, res.ID)
	if err != nil {
		return nil, fmt.Errorf("exec error %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("commit error %w", err)
	}

	return res, nil
}

type UpdateTransactionParams struct {
	ID        int64
	Status    string
	Locked    bool
	UpdatedAt time.Time
}

func (t *TransactionStorage) UpdateTransactionTurnOffLocked(ctx context.Context, p UpdateTransactionParams) error {
	var (
		querySetString strings.Builder
		queryParams    []interface{}
	)

	querySetString.WriteString("locked = ?")
	queryParams = append(queryParams, p.Locked)
	if p.Status != "" {
		querySetString.WriteString(", status = ?")
		queryParams = append(queryParams, p.Status)
	}
	queryParams = append(queryParams, p.ID)
	query := fmt.Sprintf("UPDATE transactions SET %s WHERE id = ?", querySetString.String())

	_, err := t.conn.ExecContext(ctx, t.conn.Rebind(query), queryParams)
	if err != nil {
		return fmt.Errorf("exec error %w", err)
	}

	return nil
}
