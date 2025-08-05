package httpapi

import (
	"net/http"

	"github.com/hoomanfr/harry/golib/util"
	"github.com/hoomanfr/harry/inventorysrv/internal/app"
)

type InventoryHandler struct {
	inventoryService *app.InventoryService
}

func NewInventoryHandler(inventoryService *app.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
	}
}

func (h *InventoryHandler) AddInventoryItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req AddInventoryItem
		if err := util.ShouldBindJson(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		inventoryItem := req.ToInventoryItem()
		err := h.inventoryService.AddInventoryItem(ctx, inventoryItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
