package natsclient

import (
	"context"

	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/app"
	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/infrastructure/remote"
	"github.com/hoomanfr/harry/golib/messaging/natsx"
)

func SetupRoutes(ctx context.Context, cfg *app.AppConfig, subscriber *natsx.Subscriber) error {
	inventoryClient := remote.NewFulfillmentClient(cfg)
	receivingHandler := NewReceivingHandler(inventoryClient)
	if err := subscriber.Subscribe(ctx, "wms", "receivingsrv", "receivingitem.received", receivingHandler.HandleItemReceived); err != nil {
		return err
	}

	return nil
}
