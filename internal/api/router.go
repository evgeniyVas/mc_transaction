package api

import (
	"context"
	"github.com/mc_transaction/internal/service/transaction"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Service struct {
	server *http.Server

	// processors
	transactionProc *transaction.Transaction
}

func New() *Service {
	srv := Service{}
	r := mux.NewRouter()
	r.Handle("/v1/transaction/create", http.HandlerFunc(srv.v1TransactionCreatePOST)).Methods(http.MethodPost)

	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.server = server
	return &srv
}

func (s *Service) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal("server off")
	}
}

func (s *Service) Shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}

func (s *Service) AddTransactionProc(userAuth *transaction.Transaction) {
	s.transactionProc = userAuth
}
