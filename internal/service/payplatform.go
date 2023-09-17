package service

import (
	"context"
	"fmt"
	paymodels "github.com/mc_transaction/internal/http_client/paysystem/models"
	"github.com/mc_transaction/internal/http_client/paysystem/payclient/operations"
	"github.com/mc_transaction/internal/logger"
	"net/http"
	"time"
)

type PayPlatformService struct {
	operations operations.ClientService
	httpClient *http.Client
}

func NewPayPlatformService(operations operations.ClientService) *PayPlatformService {
	return &PayPlatformService{
		operations: operations,
		httpClient: &http.Client{Timeout: 3 * time.Second},
	}
}

func (p *PayPlatformService) CreateTransaction(ctx context.Context, amount float64) (*paymodels.TransactionCreateResp, error) {
	params := &operations.PostAPITransactionsParams{
		Body:       &paymodels.TransactionCreateBody{Amount: amount},
		Context:    ctx,
		HTTPClient: p.httpClient,
	}
	//create transaction in pay platform
	resp, err := p.operations.PostAPITransactions(params)
	if err != nil {
		logger.Error("PostAPITransactions " + err.Error())
		return nil, err
	}
	return resp.Payload, nil
}

func (p *PayPlatformService) GetTransactionStatusById(ctx context.Context, payId int64) (string, error) {
	params := &operations.GetAPITransactionsIDStatusParams{
		ID:         payId,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}
	//get status transaction in pay platform
	resp, err := p.operations.GetAPITransactionsIDStatus(params)
	if err != nil {
		logger.Error("GetAPITransactionsIDStatus " + err.Error())
		return "", fmt.Errorf("payplatform error for get status transaction %w", err)
	}
	return resp.Payload.Status, nil
}

func (p *PayPlatformService) ReverseTransaction(ctx context.Context, payId int64) error {
	params := &operations.PostAPITransactionsReverseParams{
		Body:       &paymodels.TransactionReverseBody{ID: payId},
		Context:    ctx,
		HTTPClient: p.httpClient,
	}
	//reverse transaction
	_, err := p.operations.PostAPITransactionsReverse(params)
	if err != nil {
		logger.Error("PostAPITransactionsReverseParams " + err.Error())
		return fmt.Errorf("payplatform error reverse transaction %w", err)
	}
	return nil
}
