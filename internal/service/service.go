package service

import (
	"github.com/mc_transaction/internal/http_client/paysystem"
	storage "github.com/mc_transaction/internal/storage/psql"
)

type Services struct {
	BalanceService     *BalanceService
	TransactionService *TransactionService
	UserService        *UserService
	PayPlatfromService *PayPlatformService
}

type Deps struct {
	Storage           *storage.Storages
	PayPlatformClient *paysystem.PayService
}

func NewServices(deps *Deps) *Services {
	return &Services{
		BalanceService:     NewBalance(deps.Storage.Balance),
		TransactionService: NewTransaction(deps.Storage.Transaction, deps.PayPlatformClient),
		UserService:        NewUser(deps.Storage.User),
	}
}
