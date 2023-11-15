package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BalanceStorage struct {
	conn *sqlx.DB
}

func NewBalanceStorage(conn *sqlx.DB) *BalanceStorage {
	return &BalanceStorage{conn: conn}
}

var ErrBalanceNotFound = errors.New("balance not found")

type SelectBalanceParams struct {
	UserId int64
}

func (b *BalanceStorage) GetBalanceByUserID(ctx context.Context, userID int64) error {
	var res int64
	err := b.conn.GetContext(ctx, &res, "SELECT 1 FROM balance WHERE user_id=$1", userID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("select error %w", ErrBalanceNotFound)
	} else if err != nil {
		return fmt.Errorf("select error %w", err)
	}

	return nil
}

type UpdateBalanceParams struct {
	UserID            int64
	Amount            float64
	TransactionID     int64
	TransactionStatus string
}

func (b *BalanceStorage) UpdateBalanceWithUnlockTransaction(ctx context.Context, p UpdateBalanceParams) error {
	var (
		tx  *sqlx.Tx
		err error
	)
	defer rollBackTx(tx, err)

	tx, err = b.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("tx begin error: %w", err)
	}

	_, err = tx.ExecContext(ctx, queryUpdateBalance, p.Amount, p.UserID)
	if err != nil {
		return fmt.Errorf("exec error %w", err)
	}

	_, err = tx.ExecContext(ctx, queryUpdateTransactionUnlock, p.TransactionStatus, p.TransactionID)
	if err != nil {
		return fmt.Errorf("exec error %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit error %w", err)
	}

	return nil
}
