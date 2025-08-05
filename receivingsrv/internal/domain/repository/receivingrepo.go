package repository

import (
	"context"

	"github.com/hoomanfr/harry/receivingsrv/internal/domain"
	"github.com/jackc/pgx/v5"
)

type ReceivingRepo interface {
	Create(context.Context, pgx.Tx, domain.ReceivingItem) error
}
