package clients

import (
	"context"

	"github.com/hoomanfr/harry/inventoryintegrator/internal/domain"
)

type InventoryClient interface {
	AddInventoryItem(ctx context.Context, item domain.InventoryItem) error
}
