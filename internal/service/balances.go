package service

import (
	"context"
	storage "github.com/mc_transaction/internal/storage/psql"
)

type BalanceStorage interface {
	CreateBalance(ctx context.Context, fields *storage.InsertBalanceParams) int64
	UpdateBalance(ctx context.Context, fields *storage.UpdateBalanceParams) int64
}

type BalanceService struct {
	storage BalanceStorage
}

func NewBalance(storage BalanceStorage) *BalanceService {
	return &BalanceService{
		storage: storage,
	}
}

type InsertBalanceParams struct {
	Amount float64
	UserId int
}

func (t *BalanceService) CreateBalance(ctx context.Context, params InsertBalanceParams) {
	// дергается из воркера, создает баланс если у пользователя еще его нет
}

type UpdateBalanceParams struct {
	Amount float64
}

func (t *BalanceService) UpdateBalance(ctx context.Context, params UpdateBalanceParams) {
	// дергается из воркера, обновляет баланс в случае статуса SUCCESS
}
