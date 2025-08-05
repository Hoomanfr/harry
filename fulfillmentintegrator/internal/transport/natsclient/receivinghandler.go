package natsclient

import (
	"context"
	"encoding/json"

	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/domain/clients"
	"github.com/hoomanfr/harry/golib/messaging/natsx"
)

type ReceivingHandler struct {
	fulfillmentClient clients.FulfillmentClient
}

func NewReceivingHandler(fulfillmentClient clients.FulfillmentClient) *ReceivingHandler {
	return &ReceivingHandler{
		fulfillmentClient: fulfillmentClient,
	}
}

func (h *ReceivingHandler) HandleItemReceived(ctx context.Context, msg natsx.Message) error {
	var item ReceivingItemReceived
	err := json.Unmarshal(msg.Data, &item)
	if err != nil {
		return err
	}

	err = h.fulfillmentClient.FulfillBlockedOrder(ctx, item.ToFulFillmentItem())
	if err != nil {
		return err
	}

	return nil
}
