package service

import (
	"context"
	"fmt"
	paymodels "github.com/mc_transaction/internal/http_client/paysystem/models"
	"github.com/mc_transaction/internal/models"
	storage "github.com/mc_transaction/internal/storage/psql"
)

type TransactionStorage interface {
	CreateTransaction(ctx context.Context, fields *storage.InsertTransactionParams) error
}

type PayPlatformClient interface {
	CreateTransaction(ctx context.Context, amount float64) (*paymodels.TransactionCreateResp, error)
}

type TransactionService struct {
	storage     TransactionStorage
	payPlatform PayPlatformClient
}

func NewTransactionService(storage TransactionStorage, payService PayPlatformClient) *TransactionService {
	return &TransactionService{
		storage:     storage,
		payPlatform: payService,
	}
}

type InputTransactionParams struct {
	Amount float64
	UserId int64
	Token  string
}

func (t *TransactionService) CreateTransaction(ctx context.Context, params InputTransactionParams) (*models.Transaction, error) {
	//create transaction in pay platform
	resp, err := t.payPlatform.CreateTransaction(ctx, params.Amount)
	if err != nil {
		return nil, fmt.Errorf("payplatform error create transaction %w", err)
	}

	//save transaction in db
	err = t.storage.CreateTransaction(ctx, &storage.InsertTransactionParams{
		Amount: params.Amount,
		UserId: params.UserId,
		Status: "CREATED",
		PayId:  resp.PayID,
		Token:  params.Token,
	})
	if err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:      resp.PayID,
		PayLink: resp.PayURL,
	}, nil
}
