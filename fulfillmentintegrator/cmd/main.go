package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/app"
	"github.com/hoomanfr/harry/fulfillmentintegrator/internal/transport/natsclient"
	"github.com/hoomanfr/harry/golib/messaging/natsx"
)

func main() {
	ctx := context.Background()
	appConfig, err := app.NewAppConfig()
	if err != nil {
		fmt.Printf("failed to load config: %v\r\n", err)
		return
	}

	broker, err := natsx.NewBroker(*appConfig.Config, "wms", "fulfillmentintegrator")
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
		fmt.Printf("failed to setup routes: %v\r\n", err)
		return
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
}
