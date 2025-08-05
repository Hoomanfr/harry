package persistence

import (
	"context"

	"github.com/hoomanfr/harry/golib/db"
	"github.com/hoomanfr/harry/receivingsrv/internal/domain"
	"github.com/hoomanfr/harry/receivingsrv/internal/domain/repository"
	"github.com/jackc/pgx/v5"
)

type receivingRepo struct {
	pgDb *db.PgDB
}

func NewReceivingRepo(pgDb *db.PgDB) repository.ReceivingRepo {
	return &receivingRepo{pgDb: pgDb}
}

func (r *receivingRepo) Create(ctx context.Context, tx pgx.Tx, receivingItem domain.ReceivingItem) error {
	query := `INSERT INTO dbo.receiving (sku, quantity) VALUES ($1, $2)`
	_, err := tx.Exec(ctx, query, receivingItem.Sku, receivingItem.Quantity)
	if err != nil {
		return err
	}
	return nil
}
