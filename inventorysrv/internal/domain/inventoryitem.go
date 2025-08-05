package domain

import "time"

type InventoryItem struct {
	Id        int
	Sku       string
	Quantity  int
	UpdatedAt time.Time
}
