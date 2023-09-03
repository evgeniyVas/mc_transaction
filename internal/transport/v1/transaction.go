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
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, params service.InputTransactionParams) (*models.Transaction, error)
}

type UserService interface {
	GetUserByID(ctx context.Context, id int) (service.User, error)
}

type TransactionHandler struct {
	transactionService TransactionService
	userService        UserService
}

func NewTransactionHandler(transactionSer TransactionService, userServ UserService) TransactionHandler {
	return TransactionHandler{
		transactionService: transactionSer,
		userService:        userServ,
	}
}

var errAPIUnknown = errors.New("unknown service error")

const headerUser = "X-User-ID"
const headerTokenIdempodency = "X-Idempodency-Token"

func (t *TransactionHandler) TransactionCreatePOST(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	headerUserId := r.Header.Get(headerUser)
	headerToken := r.Header.Get(headerTokenIdempodency)
	headeruserIdInt, _ := strconv.Atoi(headerUserId)
	user, err := t.userService.GetUserByID(ctx, headeruserIdInt)
	if errors.Is(err, service.ErrUserNotFound) {
		rw.WriteHeader(http.StatusNotFound)
		writeErr(ctx, rw, err)
		return
	} else if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		writeErr(ctx, rw, errAPIUnknown)
		return
	}
	params, err := parseTransactionPOSTParams(ctx, rw, r)
	if err != nil {
		return
	}

	resp, err := t.transactionService.CreateTransaction(
		ctx,
		service.InputTransactionParams{
			Amount: params.Amount,
			UserId: user.Id,
			Token:  headerToken,
		},
	)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		writeErr(ctx, rw, errAPIUnknown)
		return
	}

}

func parseTransactionPOSTParams(ctx context.Context, rw http.ResponseWriter, r *http.Request) (*models.TransactionParamsBody, error) {
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

func writeErr(ctx context.Context, rw io.Writer, err error) {
	errM := &models.Error{Message: err.Error()}
	b, _ := errM.MarshalBinary()
	_, err = rw.Write(b)
	if err != nil {
		logger.Error("write: err = " + err.Error())
	}
}
