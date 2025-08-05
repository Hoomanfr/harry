package httpapi

import (
	"net/http"

	"github.com/hoomanfr/harry/golib/db"
	"github.com/hoomanfr/harry/golib/messaging/natsx"
	"github.com/hoomanfr/harry/receivingsrv/internal/app"
	"github.com/hoomanfr/harry/receivingsrv/internal/infrastructure/persistence"
)

func SetupRoutes(mux *http.ServeMux, appConfig *app.AppConfig, pgDb *db.PgDB, broker *natsx.Broker) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	receivingRepo := persistence.NewReceivingRepo(pgDb)
	receivingService := app.NewReceivingService(pgDb, broker, receivingRepo)
	receivingHandler := NewReceivingHandler(receivingService)
	mux.HandleFunc("POST /return-item", receivingHandler.ReceiveReturnedItems())
}
