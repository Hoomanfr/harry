package httpapi

import (
	"net/http"

	"github.com/hoomanfr/harry/fulfillsrv/internal/app"
	"github.com/hoomanfr/harry/fulfillsrv/internal/infrastructure/persistence"
	"github.com/hoomanfr/harry/golib/db"
	centrifugox "github.com/hoomanfr/harry/golib/messaging/centrifugo"
)

func SetupRoutes(mux *http.ServeMux, appConfig *app.AppConfig, pgDb *db.PgDB, centrifugoxBroker *centrifugox.Broker) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	fulfillmentRepo := persistence.NewFulfillmentRepo(pgDb)
	fulfillmentService := app.NewFulfillmentService(fulfillmentRepo, centrifugoxBroker)
	fulfillmentHandler := NewFulfillmentHandler(fulfillmentService)
	mux.HandleFunc("PUT /fulfillment/blocked-order", fulfillmentHandler.FulfillUnblockedOrder())
}
