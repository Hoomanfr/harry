package natsclient

import "github.com/hoomanfr/harry/inventoryintegrator/internal/domain"

type ReceivingItemReceived struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

func (r *ReceivingItemReceived) ToInventoryItem() domain.InventoryItem {
	return domain.InventoryItem{
		Sku:      r.Sku,
		Quantity: r.Quantity,
	}
}
