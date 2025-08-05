package natsclient

import (
	"context"

	"github.com/hoomanfr/harry/golib/messaging/natsx"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/app"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/infrastructure/remote"
)

func SetupRoutes(ctx context.Context, cfg *app.AppConfig, subscriber *natsx.Subscriber) error {
	inventoryClient := remote.NewInventoryClient(cfg)
	receivingHandler := NewReceivingHandler(inventoryClient)
	if err := subscriber.Subscribe(ctx, "wms", "receivingsrv", "receivingitem.received", receivingHandler.HandleItemReceived); err != nil {
		return err
	}

	return nil
}
