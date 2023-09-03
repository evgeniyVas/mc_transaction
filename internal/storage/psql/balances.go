package psql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type BalanceStorage struct {
	conn *sqlx.DB
}

func NewBalanceStorage(conn *sqlx.DB) *BalanceStorage {
	return &BalanceStorage{conn: conn}
}

type InsertBalanceParams struct {
	Amount    float64
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *BalanceStorage) CreateBalance(ctx context.Context, fields *InsertBalanceParams) int64 {
	return 123
}

type UpdateBalanceParams struct {
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *BalanceStorage) UpdateBalance(ctx context.Context, fields *UpdateBalanceParams) int64 {
	return 123
}
