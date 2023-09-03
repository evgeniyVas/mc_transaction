package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/strfmt"
	paymodels "github.com/mc_transaction/internal/http_client/paysystem/models"
	"github.com/mc_transaction/internal/http_client/paysystem/payclient"
	"github.com/mc_transaction/internal/http_client/paysystem/payclient/operations"
	"github.com/mc_transaction/internal/models"
	"github.com/mc_transaction/internal/storage/psql"
	"net/http"
	"time"
)

type TransactionStorage interface {
	CreateTransaction(ctx context.Context, fields *psql.InsertTransactionParams) (int64, error)
}

type TransactionService struct {
	storage TransactionStorage
}

func NewTransaction(storage TransactionStorage) *TransactionService {
	return &TransactionService{
		storage: storage,
	}
}

type InputTransactionParams struct {
	Amount float64
	UserId int64
	Token  string
}

var errPlatformTransactionCreate = errors.New("payplatform create transaction failed")

func (t *TransactionService) CreateTransaction(ctx context.Context, params InputTransactionParams) (*models.Transaction, error) {
	payPlatformCl := payclient.NewHTTPClient(strfmt.Default)
	httpClient := &http.Client{Timeout: 3 * time.Second}
	apiParams := &operations.PostAPITransactionsParams{
		Body:       &paymodels.TransactionCreateBody{Amount: params.Amount},
		Context:    ctx,
		HTTPClient: httpClient,
	}
	//create transaction in pay platform
	resp, err := payPlatformCl.Operations.PostAPITransactions(apiParams)
	if err != nil {
		return nil, fmt.Errorf("payplatform error create transaction %w", err)
	}
	if resp.Code() != 200 {
		return nil, errPlatformTransactionCreate
	}

	//save transaction in db
	transaction, err := t.storage.CreateTransaction(ctx, &psql.InsertTransactionParams{
		Amount:    params.Amount,
		UserId:    params.UserId,
		Status:    "CREATED",
		PayId:     resp.Payload.PayID,
		Locked:    false,
		Token:     params.Token,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	//идем в пейсистем
	//отдаем респонс модели
	//старт воркера
}
