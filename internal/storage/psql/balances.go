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

//type InsertBalanceParams struct {
//	Amount    float64
//	UserId    int
//	CreatedAt time.Time
//	UpdatedAt time.Time
//}
//
//func (s *BalanceStorage) CreateBalance(ctx context.Context, fields *InsertBalanceParams) int64 {
//	return 123
//}

type UpdateBalanceParams struct {
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BalanceStorage) UpdateBalance(ctx context.Context, fields *UpdateBalanceParams) int64 {
	return 123
}

var ErrBalanceNotFound = errors.New("balance not found")

type SelectBalanceParams struct {
	UserId int64
}

func (b *BalanceStorage) GetBalanceByUserID(ctx context.Context, fields *SelectBalanceParams) error {
	var res int64
	err := b.conn.GetContext(ctx, &res, "SELECT 1 FROM balance WHERE user_id=?", fields.UserId)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("select error %w", ErrBalanceNotFound)
	} else if err != nil {
		return fmt.Errorf("select error %w", err)
	}

	return nil
}

func (b *BalanceStorage) UpdateBalanceWithTurnLockedTransaction(ctx context.Context) error {
	return nil
}
