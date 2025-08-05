package events

import "github.com/hoomanfr/harry/receivingsrv/internal/domain"

type ReceivingItemReceived struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

func NewFromReceivingItem(receivingItem domain.ReceivingItem) ReceivingItemReceived {
	return ReceivingItemReceived{
		Sku:      receivingItem.Sku,
		Quantity: receivingItem.Quantity,
	}
}
