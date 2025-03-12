package http

import (
	"github.com/AdiPP/go-typesense-product-service/internal/app/service/sync"
	"net/http"

	"github.com/gin-gonic/gin"
)

type syncProductsHandler struct {
	productSynchronizerService *sync.Service
}

type syncProductRequest struct {
	ProductIDs []int64 `json:"product_ids"`
}

func (h *syncProductsHandler) handle(c *gin.Context) {
	var req syncProductRequest
	if c.ShouldBindBodyWithJSON(&req) != nil {
		c.String(http.StatusUnprocessableEntity, "Failed")
		return
	}

	err := h.productSynchronizerService.SyncBatch(&sync.BatchProductsParam{
		ProductIDs: req.ProductIDs,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func newSyncProductsHandler(productSynchronizerService *sync.Service) *syncProductsHandler {
	o := new(syncProductsHandler)
	o.productSynchronizerService = productSynchronizerService
	return o
}
