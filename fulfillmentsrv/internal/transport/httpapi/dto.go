package httpapi

import "github.com/hoomanfr/harry/fulfillsrv/internal/domain"

type FulfillBlockedOrder struct {
	Sku      string `json:"sku" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gt=0"`
}

func (f *FulfillBlockedOrder) ToFulFillmentItem() domain.FulFillmentItem {
	return domain.FulFillmentItem{
		Sku:      f.Sku,
		Quantity: f.Quantity,
	}
}
