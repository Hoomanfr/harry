package domain

type InventoryItem struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}
