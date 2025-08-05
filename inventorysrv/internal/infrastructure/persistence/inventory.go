package persistence

import (
	"context"

	"github.com/hoomanfr/harry/golib/db"
	"github.com/hoomanfr/harry/inventorysrv/internal/domain"
	"github.com/hoomanfr/harry/inventorysrv/internal/domain/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type inventoryRepo struct {
	pgDb *db.PgDB
}

func NewInventoryRepo(pgDb *db.PgDB) repository.InventoryRepo {
	return &inventoryRepo{pgDb: pgDb}
}

func (r *inventoryRepo) AddInventoryItem(ctx context.Context, item domain.InventoryItem) error {
	query := `INSERT INTO dbo.inventory (sku, quantity, updated_at) 
						VALUES ($1, $2, CURRENT_TIMESTAMP)
						ON CONFLICT (sku) 
						DO UPDATE SET 
								quantity = EXCLUDED.quantity,
								updated_at = CURRENT_TIMESTAMP`

	err := r.pgDb.WithConnection(ctx, func(c *pgxpool.Conn) error {
		_, err := c.Exec(ctx, query, item.Sku, item.Quantity)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
