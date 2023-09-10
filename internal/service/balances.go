package service

import (
	"context"
	"errors"
	storage "github.com/mc_transaction/internal/storage/psql"
)

type BalanceStorage interface {
	GetBalanceByUserID(ctx context.Context, fields *storage.SelectBalanceParams) error
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

var ErrBalanceNotFound = errors.New("balance not found")

func (b *BalanceService) CheckBalanceByUserID(ctx context.Context, userId int64) error {
	err := b.storage.GetBalanceByUserID(ctx, &storage.SelectBalanceParams{
		UserId: userId,
	})
	if errors.Is(err, storage.ErrBalanceNotFound) {
		return ErrBalanceNotFound
	} else if err != nil {
		return err
	}
	return nil
}

type UpdateBalanceParams struct {
	Amount float64
	UserId int64
}

func (b *BalanceService) UpdateBalance(ctx context.Context, params UpdateBalanceParams) {
	// дергается из воркера, обновляет баланс в случае статуса SUCCESS
}
