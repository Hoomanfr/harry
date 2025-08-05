package config

type Config struct {
	DbUrl               string `env:"DB_URL"`
	DbMaxConns          int32  `env:"DB_MAX_CONNS" envDefault:"20"`
	DbMinConns          int32  `env:"DB_MIN_CONNS" envDefault:"2"`
	DbMaxConnIdleTime   int    `env:"DB_MAX_CONN_IDLE_TIME" envDefault:"15"`
	DbMaxConnLifetime   int    `env:"DB_MAX_CONN_LIFETIME" envDefault:"30"`
	DbHealthCheckPeriod int    `env:"DB_HEALTH_CHECK_PERIOD" envDefault:"1"`
	NatsUrls            string `env:"NATS_URLS"`
	NatsMaxReconnects   int    `env:"NATS_MAX_RECONNECTS" envDefault:"10"`
	NatsReconnectWait   int    `env:"NATS_RECONNECT_WAIT" envDefault:"1"`
	CentrifugoUrl       string `env:"CENTRIFUGO_URL"`
	CentrifugoToken     string `env:"CENTRIFUGO_TOKEN"`
}
