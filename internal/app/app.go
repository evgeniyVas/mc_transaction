// Package app предоставляет основные структуры для работы с приложением.
// Пакет позволяет получить конфигурации приложения и определить в каком режиме оно запущено.
package app

import (
	"context"
	"fmt"
	"github.com/mc_transaction/internal/api"
	"github.com/mc_transaction/internal/service/transaction"
	"github.com/mc_transaction/internal/storage/psql"
)

const (
	ServiceName = "mc_transaction"
)

type App struct {
	server  *api.Service
	storage *psql.Storage
}

func New() (*App, error) {
	cfg, err := NewConfig(ServiceName)
	if err != nil {
		return nil, fmt.Errorf("config initialization error: (%w)", err)
	}

	srv := api.New()

	storage, err := psql.New(
		cfg.Postgres.Dsn,
		cfg.Postgres.PingTimeout,
	)
	if err != nil {
		return nil, fmt.Errorf("storage initialization error: (%w)", err)
	}

	srv.AddTransactionProc(transaction.New(storage))

	return &App{
		storage: storage,
		server:  srv,
	}, nil
}

func (a App) Start() {
	a.server.Start()
}

func (a App) Stop(ctx context.Context) {
	a.server.Shutdown(ctx)
	a.storage.Close()
}
