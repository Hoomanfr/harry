package httpapi

import "github.com/hoomanfr/harry/inventorysrv/internal/domain"

type AddInventoryItem struct {
	Sku      string `json:"sku" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gt=0"`
}

func (a *AddInventoryItem) ToInventoryItem() domain.InventoryItem {
	return domain.InventoryItem{
		Sku:      a.Sku,
		Quantity: a.Quantity,
	}
}
