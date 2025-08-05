package events

type OrderUnblocked struct {
	OrderId  int    `json:"order_id"`
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}
