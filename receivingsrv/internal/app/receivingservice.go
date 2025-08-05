package app

import (
	"context"

	"github.com/hoomanfr/harry/golib/db"
	"github.com/hoomanfr/harry/golib/messaging/natsx"
	"github.com/hoomanfr/harry/receivingsrv/internal/domain"
	"github.com/hoomanfr/harry/receivingsrv/internal/domain/events"
	"github.com/hoomanfr/harry/receivingsrv/internal/domain/repository"
	"github.com/jackc/pgx/v5"
)

type ReceivingService struct {
	pgDb          *db.PgDB
	broker        *natsx.Broker
	receivingRepo repository.ReceivingRepo
}

func NewReceivingService(pgDb *db.PgDB, broker *natsx.Broker, receivingRepo repository.ReceivingRepo) *ReceivingService {
	return &ReceivingService{
		pgDb:          pgDb,
		broker:        broker,
		receivingRepo: receivingRepo,
	}
}

func (s *ReceivingService) ReceiveReturnedItems(ctx context.Context, receivingItem domain.ReceivingItem) error {
	// Double write issue, using transaction can mitigate it but not completely solve it.
	err := s.pgDb.WithTransaction(ctx, func(tx pgx.Tx) error {
		err := s.receivingRepo.Create(ctx, tx, receivingItem)
		if err != nil {
			return err
		}
		err = s.broker.Publish("receivingitem.received", events.NewFromReceivingItem(receivingItem))
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
