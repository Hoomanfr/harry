package persistence

import (
	"context"

	"github.com/hoomanfr/harry/fulfillsrv/internal/domain"
	"github.com/hoomanfr/harry/fulfillsrv/internal/domain/repository"
	"github.com/hoomanfr/harry/golib/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type fulfillmentRepo struct {
	pgDb *db.PgDB
}

func NewFulfillmentRepo(pgDb *db.PgDB) repository.FulfillmentRepository {
	return &fulfillmentRepo{pgDb: pgDb}
}

func (r *fulfillmentRepo) GetPriorityBlockedOrderWithQuantity(ctx context.Context, sku string, quantity int, status string) (*domain.FulFillmentItem, error) {
	query := `SELECT id, order_id, sku, quantity, status, fulfilled_at FROM dbo.fulfillment WHERE sku = $1 AND quantity = $2 AND status = $3 ORDER BY fulfilled_at ASC LIMIT 1`

	var item domain.FulFillmentItem
	err := r.pgDb.WithConnection(ctx, func(c *pgxpool.Conn) error {
		err := c.QueryRow(ctx, query, sku, quantity, status).Scan(
			&item.Id,
			&item.OrderId,
			&item.Sku,
			&item.Quantity,
			&item.Status,
			&item.FulfilledAt,
		)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}
