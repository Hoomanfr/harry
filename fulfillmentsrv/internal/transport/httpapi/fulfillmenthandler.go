package httpapi

import (
	"net/http"

	"github.com/hoomanfr/harry/fulfillsrv/internal/app"
	"github.com/hoomanfr/harry/golib/util"
)

type FulfillmentHandler struct {
	FulfillmentService *app.FulfillmentService
}

func NewFulfillmentHandler(fulfillmentService *app.FulfillmentService) *FulfillmentHandler {
	return &FulfillmentHandler{
		FulfillmentService: fulfillmentService,
	}
}

func (h *FulfillmentHandler) FulfillUnblockedOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req FulfillBlockedOrder
		if err := util.ShouldBindJson(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fulfillmentItem := req.ToFulFillmentItem()
		err := h.FulfillmentService.FulfillBlockedOrder(ctx, &fulfillmentItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
