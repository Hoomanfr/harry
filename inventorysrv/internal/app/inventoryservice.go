package app

import (
	"context"

	"github.com/hoomanfr/harry/inventorysrv/internal/domain"
	"github.com/hoomanfr/harry/inventorysrv/internal/domain/repository"
)

type InventoryService struct {
	inventoryRepo repository.InventoryRepo
}

func NewInventoryService(inventoryRepo repository.InventoryRepo) *InventoryService {
	return &InventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (s *InventoryService) AddInventoryItem(ctx context.Context, item domain.InventoryItem) error {
	return s.inventoryRepo.AddInventoryItem(ctx, item)
}
