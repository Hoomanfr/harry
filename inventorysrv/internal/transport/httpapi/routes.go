package httpapi

import (
	"net/http"

	"github.com/hoomanfr/harry/golib/db"
	"github.com/hoomanfr/harry/inventorysrv/internal/app"
	"github.com/hoomanfr/harry/inventorysrv/internal/infrastructure/persistence"
)

func SetupRoutes(mux *http.ServeMux, appConfig *app.AppConfig, pgDb *db.PgDB) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	inventoryRepo := persistence.NewInventoryRepo(pgDb)
	inventoryService := app.NewInventoryService(inventoryRepo)
	inventoryHandler := NewInventoryHandler(inventoryService)
	mux.HandleFunc("POST /inventory", inventoryHandler.AddInventoryItem())
}
