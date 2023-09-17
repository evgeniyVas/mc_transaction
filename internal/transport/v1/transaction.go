package v1

import (
	"context"
	"errors"
	"github.com/mc_transaction/internal/logger"
	"github.com/mc_transaction/internal/models"
	"github.com/mc_transaction/internal/service"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, params service.InputTransactionParams) (*models.Transaction, error)
}

type BalanceService interface {
	CheckBalanceByUserID(ctx context.Context, userId int64) error
}

type TransactionHandler struct {
	transactionService TransactionService
	balanceService     BalanceService
}

func NewTransactionHandler(transactionServ TransactionService, balanceServ BalanceService) TransactionHandler {
	return TransactionHandler{
		transactionService: transactionServ,
		balanceService:     balanceServ,
	}
}

var errAPIUnknown = errors.New("unknown service error")

const headerUser = "X-User-ID"
const headerTokenIdempodency = "X-Idempodency-Token"

func (t *TransactionHandler) TransactionCreatePOST(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	headerUserId := strings.TrimSpace(r.Header.Get(headerUser))
	headerToken := strings.TrimSpace(r.Header.Get(headerTokenIdempodency))
	userId, _ := strconv.ParseInt(headerUserId, 10, 64)

	err := t.balanceService.CheckBalanceByUserID(ctx, userId)
	if errors.Is(err, service.ErrBalanceNotFound) {
		rw.WriteHeader(http.StatusNotFound)
		writeErr(rw, err)
		return
	} else if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		writeErr(rw, errAPIUnknown)
		return
	}
	params, err := parseTransactionPOSTParams(rw, r)
	if err != nil {
		return
	}

	resp, err := t.transactionService.CreateTransaction(
		ctx,
		service.InputTransactionParams{
			Amount: params.Amount,
			UserId: userId,
			Token:  headerToken,
		},
	)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		writeErr(rw, errAPIUnknown)
		return
	}

	bytesResp, err := resp.MarshalBinary()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		writeErr(rw, errAPIUnknown)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	_, err = rw.Write(bytesResp)
	if err != nil {
		logger.Error("TransactionCreatePOST: err = " + err.Error())
		return
	}
}

func parseTransactionPOSTParams(rw http.ResponseWriter, r *http.Request) (*models.TransactionParamsBody, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		logger.Info("v1TransactionPOST: err = " + err.Error())
		return nil, err
	}

	input := &models.TransactionParamsBody{}
	err = input.UnmarshalBinary(body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		logger.Info("v1TransactionPOST: err = " + err.Error())
		return nil, err
	}

	return input, nil
}

func writeErr(rw io.Writer, err error) {
	errM := &models.Error{Message: err.Error()}
	b, _ := errM.MarshalBinary()
	_, err = rw.Write(b)
	if err != nil {
		logger.Error("write: err = " + err.Error())
	}
}
