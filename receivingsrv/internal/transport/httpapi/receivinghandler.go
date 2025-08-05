package httpapi

import (
	"net/http"

	"github.com/hoomanfr/harry/golib/util"
	"github.com/hoomanfr/harry/receivingsrv/internal/app"
)

type ReceivingHandler struct {
	receivingService *app.ReceivingService
}

func NewReceivingHandler(receivingService *app.ReceivingService) *ReceivingHandler {
	return &ReceivingHandler{
		receivingService: receivingService,
	}
}

func (h *ReceivingHandler) ReceiveReturnedItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req ReceiveReturnedItem
		if err := util.ShouldBindJson(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		receivingItem := req.ToReceivingItem()
		err := h.receivingService.ReceiveReturnedItems(ctx, receivingItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
