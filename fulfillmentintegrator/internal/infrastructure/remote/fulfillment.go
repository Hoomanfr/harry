package remote

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/app"
	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/domain"
	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/domain/clients"
	"github.com/hoomanfr/harry/golib/util"
)

type fulfillmentClient struct {
	cfg        *app.AppConfig
	httpClient http.Client
}

func NewFulfillmentClient(cfg *app.AppConfig) clients.FulfillmentClient {
	httpClient := util.NewHttpClient(
		util.WithTimeout(cfg.HttpClientTimeout),
	)
	return &fulfillmentClient{
		cfg:        cfg,
		httpClient: httpClient,
	}
}

func (f *fulfillmentClient) FulfillBlockedOrder(ctx context.Context, item domain.FulfillmentItem) error {
	url := f.cfg.FulfillmentServiceURL + "/fulfillment/blocked-order"
	bodyJson, err := json.Marshal(item)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(bodyJson))
	if err != nil {
		return err
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("failed to close response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		if body, err := io.ReadAll(resp.Body); err == nil {
			fmt.Printf("response error body: %s\n", body)
		} else {
			fmt.Printf("failed to read response error body: %v\n", err)
		}
		return domain.ErrOrderNotFulfilled
	}

	return nil
}
