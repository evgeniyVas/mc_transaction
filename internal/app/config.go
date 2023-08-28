package app

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres *Postgres
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
