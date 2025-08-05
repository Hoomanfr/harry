package app

import (
	"github.com/caarlos0/env/v11"
	"github.com/hoomanfr/harry/golib/config"
)

type AppConfig struct {
	*config.Config
	InventoryServiceURL string `env:"INVENTORY_SERVICE_URL"`
	HttpClientTimeout   int    `env:"HTTP_CLIENT_TIMEOUT" envDefault:"10"`
}

func NewAppConfig() (*AppConfig, error) {
	cfg := AppConfig{
		Config: &config.Config{},
	}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
