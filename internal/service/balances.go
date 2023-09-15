package service

import (
	"context"
	"errors"
	storage "github.com/mc_transaction/internal/storage/psql"
)

type BalanceStorage interface {
	GetBalanceByUserID(ctx context.Context, userID int64) error
}

type BalanceService struct {
	storage BalanceStorage
}

func NewBalanceService(storage BalanceStorage) *BalanceService {
	return &BalanceService{
		storage: storage,
	}
}

var ErrBalanceNotFound = errors.New("balance not found")

func (b *BalanceService) CheckBalanceByUserID(ctx context.Context, userId int64) error {
	err := b.storage.GetBalanceByUserID(ctx, userId)
	if errors.Is(err, storage.ErrBalanceNotFound) {
		return ErrBalanceNotFound
	} else if err != nil {
		return err
	}
	return nil
}
