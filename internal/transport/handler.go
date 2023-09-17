package transport

import (
	"github.com/gorilla/mux"
	"github.com/mc_transaction/internal/service"
	v1 "github.com/mc_transaction/internal/transport/v1"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter() http.Handler {
	r := mux.NewRouter()

	transactionHandler := v1.NewTransactionHandler(h.services.TransactionService, h.services.BalanceService)
	r.Handle("/v1/transaction/create", http.HandlerFunc(transactionHandler.TransactionCreatePOST)).Methods(http.MethodPost)
	return r
}
