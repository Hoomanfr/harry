package repository

import (
	"context"

	"github.com/hoomanfr/harry/fulfillsrv/internal/domain"
)

type FulfillmentRepository interface {
	GetPriorityBlockedOrderWithQuantity(context.Context, string, int, string) (*domain.FulFillmentItem, error)
}
