package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
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
	err := b.conn.GetContext(ctx, &res, "SELECT 1 FROM balance WHERE user_id=?", userID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("select error %w", ErrBalanceNotFound)
	} else if err != nil {
		return fmt.Errorf("select error %w", err)
	}

	return nil
}

type UpdateBalanceParams struct {
	UserID    int64
	Amount    float64
	UpdatedAt time.Time
}

func (b *BalanceStorage) UpdateBalance(ctx context.Context, p UpdateBalanceParams) error {
	_, err := b.conn.ExecContext(ctx, queryUpdateBalance, p.Amount, p.UserID)
	if err != nil {
		return fmt.Errorf("exec error %w", err)
	}

	return nil
}
