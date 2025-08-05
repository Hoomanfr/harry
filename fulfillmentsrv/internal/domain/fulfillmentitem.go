package domain

import "time"

const (
	FulfillmentStatusPending   = "pending"
	FulfillmentStatusFulfilled = "fulfilled"
	FulfillmentStatusBlocked   = "blocked"
)

type FulFillmentItem struct {
	Id          int
	OrderId     int
	Sku         string
	Quantity    int
	Status      string
	FulfilledAt time.Time
}
