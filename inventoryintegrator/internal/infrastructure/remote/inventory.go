package remote

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hoomanfr/harry/golib/util"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/app"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/domain"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/domain/clients"
)

type inventoryClient struct {
	cfg        *app.AppConfig
	httpClient http.Client
}

func NewInventoryClient(cfg *app.AppConfig) clients.InventoryClient {
	httpClient := util.NewHttpClient(
		util.WithTimeout(cfg.HttpClientTimeout),
	)
	return &inventoryClient{
		cfg:        cfg,
		httpClient: httpClient,
	}
}

func (c *inventoryClient) AddInventoryItem(ctx context.Context, item domain.InventoryItem) error {
	url := c.cfg.InventoryServiceURL + "/inventory"

	bodyBytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("failed to close response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("failed to read response body: %v\n", err)
		} else {
			fmt.Printf("response body: %s\n", body)
		}
		return domain.ErrInventoryItemNotAdded
	}

	return nil
}
