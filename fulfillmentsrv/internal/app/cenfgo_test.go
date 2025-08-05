package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/hoomanfr/harry/fulfillsrv/internal/domain/events"
	"github.com/hoomanfr/harry/golib/config"
	centrifugox "github.com/hoomanfr/harry/golib/messaging/centrifugo"
)

func Test_Cenfgo(t *testing.T) {
	ctx := context.Background()
	appConfig := &AppConfig{
		Config: &config.Config{
			CentrifugoUrl:   "ws://localhost:8000/connection/websocket?cf_ws_frame_ping_pong=true",
			CentrifugoToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM3MjIiLCJleHAiOjE3NTQ4NzIwOTcsImlhdCI6MTc1NDI2NzI5N30.Uqp-YwEO9ZvPQNW8HBVsRARq_c7hnEmPmW_UO6RHzfE",
		},
	}
	centrifugox, err := centrifugox.NewBroker(appConfig.Config)
	if err != nil {
		fmt.Printf("failed to create centrifugo broker: %v\r\n", err)
		return
	}

	if err = centrifugox.Connect(); err != nil {
		fmt.Printf("failed to connect to centrifugo: %v\r\n", err)
		return
	}
	orderUnblocked := &events.OrderUnblocked{
		OrderId:  2222,
		Sku:      "sku-1234",
		Quantity: 12,
	}
	err = centrifugox.Publish(ctx, "fulfillment.order_unblocked", orderUnblocked)
	if err != nil {
		fmt.Println("failed to publish order unblocked event:", err)
	}
}
