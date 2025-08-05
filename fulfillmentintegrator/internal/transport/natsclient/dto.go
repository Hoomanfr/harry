package natsclient

import (
	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/domain"
)

type ReceivingItemReceived struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

func (r *ReceivingItemReceived) ToFulFillmentItem() domain.FulfillmentItem {
	return domain.FulfillmentItem{
		Sku:      r.Sku,
		Quantity: r.Quantity,
	}
}
