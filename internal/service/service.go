package service

import storage "github.com/mc_transaction/internal/storage/psql"

type Services struct {
	BalanceService     *BalanceService
	TransactionService *TransactionService
	UserService        *UserService
}

func NewServices(storages *storage.Storages) *Services {
	return &Services{
		BalanceService:     NewBalance(storages.Balance),
		TransactionService: NewTransaction(storages.Transaction),
		UserService:        NewUser(storages.User),
	}
}
