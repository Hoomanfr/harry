package app

import (
	"github.com/caarlos0/env/v11"
	"github.com/hoomanfr/harry/golib/config"
)

type AppConfig struct {
	*config.Config
	Port int `env:"PORT" envDefault:"8082"`
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
