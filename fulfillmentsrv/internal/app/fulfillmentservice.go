package app

import (
	"context"

	"github.com/hoomanfr/harry/fulfillsrv/internal/domain"
	"github.com/hoomanfr/harry/fulfillsrv/internal/domain/events"
	"github.com/hoomanfr/harry/fulfillsrv/internal/domain/repository"
	centrifugox "github.com/hoomanfr/harry/golib/messaging/centrifugo"
)

type FulfillmentService struct {
	fulfillmentRepo  repository.FulfillmentRepository
	centrifugoBroker *centrifugox.Broker
}

func NewFulfillmentService(fulfillmentRepo repository.FulfillmentRepository, centrifugoBroker *centrifugox.Broker) *FulfillmentService {
	return &FulfillmentService{
		fulfillmentRepo:  fulfillmentRepo,
		centrifugoBroker: centrifugoBroker,
	}
}

func (s *FulfillmentService) FulfillBlockedOrder(ctx context.Context, fulFillmentItem *domain.FulFillmentItem) error {
	var err error
	fulFillmentItem, err = s.fulfillmentRepo.GetPriorityBlockedOrderWithQuantity(ctx, fulFillmentItem.Sku, fulFillmentItem.Quantity, domain.FulfillmentStatusBlocked)
	if err != nil {
		return err
	}
	if fulFillmentItem == nil {
		return nil
	}
	orderUnblocked := &events.OrderUnblocked{
		OrderId:  fulFillmentItem.OrderId,
		Sku:      fulFillmentItem.Sku,
		Quantity: fulFillmentItem.Quantity,
	}
	err = s.centrifugoBroker.Publish(ctx, "fulfillment.order_unblocked", orderUnblocked)
	if err != nil {
		return err
	}
	return nil
}
