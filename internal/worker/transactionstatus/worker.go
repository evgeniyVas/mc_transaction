package transactionstatus

import (
	"context"
	"github.com/mc_transaction/pkg/worker"
	"time"
)

type TransactionStorage interface {
	SelectTransactionWithLock(ctx context.Context) (int64, error)
	UpdateTransactionStatus(ctx context.Context) error
}

type BalanceStorage interface {
	UpdateBalanceWithTurnLockedTransaction(ctx context.Context) error
}

type PayPlatformClient interface {
	GetTransactionStatusById(ctx context.Context, payId int64) (string, error)
	ReverseTransaction(ctx context.Context, payId int64) error
}

type Cfg struct {
	Period         time.Duration `default:"2s"`
	Count          int           `default:"100"`
	LockedDuration time.Duration `default:"60s"`
}

type TransactionWorker struct {
	tStorage TransactionStorage
	bStorage BalanceStorage
	pClient  PayPlatformClient
	workers  *worker.Workers
	cfg      Cfg
}

const workerName = "statusTransactionWorker"

func NewTransactionWorker(cfg Cfg, pClient PayPlatformClient, tStorage TransactionStorage, bStorage BalanceStorage) *TransactionWorker {
	transactionW := TransactionWorker{
		tStorage: tStorage,
		bStorage: bStorage,
		pClient:  pClient,
		cfg:      cfg,
	}
	transactionW.workers = worker.New(workerName, cfg.Count, cfg.Period, transactionW.work)
	return &transactionW
}

func (s *TransactionWorker) Start(ctx context.Context) {
	s.workers.Start(ctx)
}

func (s *TransactionWorker) Close() {
	s.workers.Close()
}

func (s *TransactionWorker) FuncTrigger() func() {
	return s.workers.Trigger
}

func (s *TransactionWorker) work(ctx context.Context) {
	ctx, cancelFunc := context.WithTimeout(ctx, s.cfg.LockedDuration)
	defer cancelFunc()

	//logic
	s.workers.Trigger()
}
