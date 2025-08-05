package main

import (
	"fmt"
	"net/http"

	"github.com/hoomanfr/harry/fulfillsrv/internal/app"
	"github.com/hoomanfr/harry/fulfillsrv/internal/transport/httpapi"
	"github.com/hoomanfr/harry/golib/db"
	centrifugox "github.com/hoomanfr/harry/golib/messaging/centrifugo"
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

	centrifugox, err := centrifugox.NewBroker(appConfig.Config)
	if err != nil {
		fmt.Printf("failed to create centrifugo broker: %v\r\n", err)
		return
	}

	if err = centrifugox.Connect(); err != nil {
		fmt.Printf("failed to connect to centrifugo: %v\r\n", err)
		return
	}
	fmt.Println("app config is: ", appConfig.Config)
	mux := http.NewServeMux()
	httpapi.SetupRoutes(mux, appConfig, pgDb, centrifugox)
	err = http.ListenAndServe(":8082", mux)
	if err != nil {
		fmt.Printf("failed to start server: %v\r\n", err)
	}
}
