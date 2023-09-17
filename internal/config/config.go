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
	Port         string        `default:"8085"`
	ReadTimeout  time.Duration `default:"6s"`
	WriteTimeout time.Duration `default:"6s"`
}

type Postgres struct {
	Dsn         string        `default:"host=0.0.0.0 port=5429 user=postgres password=postgres dbname=transactions sslmode=disable binary_parameters=yes"`
	PingTimeout time.Duration `default:"60s"`
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
