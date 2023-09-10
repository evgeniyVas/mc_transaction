package config

import (
	"github.com/mc_transaction/internal/worker/transactionstatus"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres                *Postgres
	HttpServer              *HttpServer
	TransactionStatusWorker transactionstatus.Cfg
}

type HttpServer struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Postgres struct {
	Dsn         string
	PingTimeout time.Duration
}

// NewConfig возвращает ссылку на объект конфигурации.
func NewConfig(serviceName string) (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(serviceName, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
