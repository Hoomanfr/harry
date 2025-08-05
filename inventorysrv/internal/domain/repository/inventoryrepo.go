package repository

import (
	"context"

	"github.com/hoomanfr/harry/inventorysrv/internal/domain"
)

type InventoryRepo interface {
	AddInventoryItem(context.Context, domain.InventoryItem) error
}
