package transactionstatus

import (
	"context"
	"errors"
	"github.com/mc_transaction/internal/logger"
	storage "github.com/mc_transaction/internal/storage/psql"
	"github.com/mc_transaction/pkg/worker"
	"time"
)

type TransactionStorage interface {
	SelectTransactionWithLock(ctx context.Context) (*storage.Transaction, error)
	UpdateTransactionTurnOffLocked(ctx context.Context, params storage.UpdateTransactionParams) error
}

type BalanceStorage interface {
	UpdateBalanceWithUnlockTransaction(ctx context.Context, fields storage.UpdateBalanceParams) error
}

type PayPlatformClient interface {
	GetTransactionStatusById(ctx context.Context, payId int64) (string, error)
	ReverseTransaction(ctx context.Context, payId int64) error
}

type Cfg struct {
	Period          time.Duration `default:"2s"`
	Count           int           `default:"40"`
	ReverseDuration time.Duration `default:"15m"`
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

func (t *TransactionWorker) Start(ctx context.Context) {
	t.workers.Start(ctx)
}

func (t *TransactionWorker) Close() {
	t.workers.Close()
}

func (t *TransactionWorker) FuncTrigger() func() {
	return t.workers.Trigger
}

func (t *TransactionWorker) work(ctx context.Context) {
	const logPrefix = "transactionStatusWorker:"
	transaction, err := t.tStorage.SelectTransactionWithLock(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrTransactionNotFound) {
			return
		}
		logger.Error(logPrefix + " error to select transaction " + err.Error())
		return
	}
	tParams := storage.UpdateTransactionParams{
		ID: transaction.ID,
	}

	useUnlockTrans := true
	defer func() {
		if useUnlockTrans {
			err = t.tStorage.UpdateTransactionTurnOffLocked(ctx, tParams)
			if err != nil {
				logger.Error(logPrefix + " error update transaction - " + err.Error())
			}
		}
	}()

	status, err := t.pClient.GetTransactionStatusById(ctx, transaction.PayId)
	if err != nil {
		logger.Error(logPrefix + " error payplatform get status endpoint - " + err.Error())
		return
	}

	if status == "CREATED" {
		if time.Now().Sub(transaction.CreatedAt.Add(t.cfg.ReverseDuration)) > 0 {
			err := t.pClient.ReverseTransaction(ctx, transaction.PayId)
			if err != nil {
				logger.Error(logPrefix + " error payplatform reverse endpoint - " + err.Error())
				return
			}
			tParams.Status = "reverse"
		}
	} else if status == "SUCCESS" {
		err := t.bStorage.UpdateBalanceWithUnlockTransaction(ctx, storage.UpdateBalanceParams{
			UserID:            transaction.UserId,
			Amount:            transaction.Amount,
			TransactionID:     transaction.ID,
			TransactionStatus: "success",
		})
		if err != nil {
			logger.Error(logPrefix + " error update balance - " + err.Error())
			return
		}
		useUnlockTrans = false
	} else if status == "NEED_TO_REVERSE" {
		err := t.pClient.ReverseTransaction(ctx, transaction.PayId)
		if err != nil {
			logger.Error(logPrefix + " error payplatform reverse endpoint - " + err.Error())
			return
		}
		tParams.Status = "reverse"
	}

	t.workers.Trigger()
}
