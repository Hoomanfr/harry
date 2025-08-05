package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoomanfr/harry/golib/messaging/natsx"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/app"
	"github.com/hoomanfr/harry/inventoryintegrator/internal/transport/natsclient"
)

func main() {
	ctx := context.Background()
	appConfig, err := app.NewAppConfig()
	if err != nil {
		fmt.Printf("failed to load config: %v\r\n", err)
		return
	}
	broker, err := natsx.NewBroker(*appConfig.Config, "wms", "inventoryintegrator")
	if err != nil {
		fmt.Printf("failed to create nats broker: %v\r\n", err)
		return
	}

	err = broker.Connect()
	if err != nil {
		fmt.Printf("failed to connect to nats broker: %v\r\n", err)
		return
	}

	subscriber := natsx.NewSubscriber(broker)

	err = natsclient.SetupRoutes(ctx, appConfig, subscriber)

	if err != nil {
		fmt.Printf("failed to setup routes: nats:%s - err: %v\r\n", appConfig.NatsUrls, err)
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
}
