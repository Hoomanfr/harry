package clients

import (
	"context"

	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/domain"
)

type FulfillmentClient interface {
	FulfillBlockedOrder(ctx context.Context, item domain.FulfillmentItem) error
}
