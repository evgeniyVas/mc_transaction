package main

import (
	"context"
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/mc_transaction/internal/config"
	"github.com/mc_transaction/internal/http_client/paysystem"
	"github.com/mc_transaction/internal/http_client/paysystem/payclient"
	"github.com/mc_transaction/internal/logger"
	"github.com/mc_transaction/internal/service"
	storage "github.com/mc_transaction/internal/storage/psql"
	"github.com/mc_transaction/internal/worker/transactionstatus"

	"github.com/mc_transaction/internal/transport"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ServiceName = "mc_transaction"
)

func main() {
	cfg, err := config.NewConfig(ServiceName)
	if err != nil {
		panic("config initialization error: " + err.Error())
	}

	logger.NewLogger()

	storages, err := storage.New(
		cfg.Postgres.Dsn,
		cfg.Postgres.PingTimeout,
	)
	if err != nil {
		panic("storage initialization error: " + err.Error())
	}

	payPlatformClient := paysystem.NewPayPlatformService(payclient.NewHTTPClient(strfmt.Default).Operations)
	transactionWorker := transactionstatus.NewTransactionWorker(
		cfg.TransactionStatusWorker,
		payPlatformClient,
		storages.Transaction,
		storages.Balance,
	)
	transactionWorker.Start(context.Background())
	defer transactionWorker.Close()

	services := service.NewServices(&service.Deps{
		Storage:           storages,
		PayPlatformClient: payPlatformClient,
	})
	handler := transport.NewHandler(services)
	server := transport.NewServer(cfg, handler.InitRouter())

	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			panic("error occurred while running http server: " + err.Error())
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	err = server.Stop(ctx)
	if err != nil {
		logger.Error("failed to stop server:" + err.Error())
	}

	err = storages.Close()
	if err != nil {
		logger.Error("failed to close storage:" + err.Error())
	}
}
