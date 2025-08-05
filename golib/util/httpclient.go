package util

import (
	"net/http"
	"time"
)

type HttpClientOptions func(*http.Client)

func NewHttpClient(opts ...HttpClientOptions) http.Client {
	client := http.Client{}
	for _, opt := range opts {
		opt(&client)
	}
	return client
}

func WithTimeout(timeout int) HttpClientOptions {
	return func(c *http.Client) {
		c.Timeout = time.Duration(timeout) * time.Second
	}
}
