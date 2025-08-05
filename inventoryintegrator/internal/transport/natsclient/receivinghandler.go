package natsclient

import (
	"context"
	"encoding/json"

	"github.com/hoomanfr/harry/golib/messaging/natsx"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/domain/clients"
)

type ReceivingHandler struct {
	inventoryClient clients.InventoryClient
}

func NewReceivingHandler(inventoryClient clients.InventoryClient) *ReceivingHandler {
	return &ReceivingHandler{
		inventoryClient: inventoryClient,
	}
}

func (h *ReceivingHandler) HandleItemReceived(ctx context.Context, msg natsx.Message) error {
	var item ReceivingItemReceived
	err := json.Unmarshal(msg.Data, &item)
	if err != nil {
		return err
	}

	err = h.inventoryClient.AddInventoryItem(ctx, item.ToInventoryItem())
	if err != nil {
		return err
	}

	return nil
}
