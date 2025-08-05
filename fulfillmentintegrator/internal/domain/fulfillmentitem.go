package domain

type FulfillmentItem struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}
