package httpapi

import "github.com/hoomanfr/harry/receivingsrv/internal/domain"

type ReceiveReturnedItem struct {
	Sku      string `json:"sku" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gt=0"`
}

func (r *ReceiveReturnedItem) ToReceivingItem() domain.ReceivingItem {
	return domain.ReceivingItem{
		Sku:      r.Sku,
		Quantity: r.Quantity,
	}
}
