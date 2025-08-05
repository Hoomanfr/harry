package main

import (
	"fmt"
	"net/http"

	"github.com/hoomanfr/harry/golib/db"
	"github.com/hoomanfr/harry/inventorysrv/internal/app"
	"github.com/hoomanfr/harry/inventorysrv/internal/transport/httpapi"
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
	mux := http.NewServeMux()
	httpapi.SetupRoutes(mux, appConfig, pgDb)
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Printf("failed to start server: %v\r\n", err)
	}
}
