package main

import (
	"fmt"
	"net/http"

	"github.com/hoomanfr/harry/golib/db"
	"github.com/hoomanfr/harry/golib/messaging/natsx"
	"github.com/hoomanfr/harry/receivingsrv/internal/app"
	"github.com/hoomanfr/harry/receivingsrv/internal/transport/httpapi"
)

func main() {
	appConfig, err := app.NewAppConfig()
	if err != nil {
		fmt.Printf("failed to load config: %v\r\n", err)
		return
	}
	pgDb, err := db.NewPostgresConnection(*appConfig.Config)
	if err != nil {
		fmt.Printf("failed to connect to database: %v\r\n", err)
		return
	}
	broker, err := natsx.NewBroker(*appConfig.Config, "wms", "receivingsrv")
	if err != nil {
		fmt.Printf("failed to create nats broker: %v\r\n", err)
		return
	}
	err = broker.Connect()
	if err != nil {
		fmt.Printf("failed to connect to nats broker: %v\r\n", err)
		return
	}
	mux := http.NewServeMux()
	httpapi.SetupRoutes(mux, appConfig, pgDb, broker)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("failed to start server: %v\r\n", err)
	}
}
