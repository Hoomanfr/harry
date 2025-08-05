package domain

import "time"

type ReceivingItem struct {
	Id         int
	Sku        string
	Quantity   int
	ReceivedAt time.Time
}
