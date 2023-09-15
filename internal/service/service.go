package service

import (
	"github.com/mc_transaction/internal/http_client/paysystem"
	storage "github.com/mc_transaction/internal/storage/psql"
)

type Services struct {
	BalanceService     *BalanceService
	TransactionService *TransactionService
}

type Deps struct {
	Storage           *storage.Storages
	PayPlatformClient *paysystem.PayService
}

func NewServices(deps *Deps) *Services {
	return &Services{
		BalanceService:     NewBalanceService(deps.Storage.BalanceStorage),
		TransactionService: NewTransactionService(deps.Storage.TransactionStorage, deps.PayPlatformClient),
	}
}
