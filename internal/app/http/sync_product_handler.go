package http

import (
	"net/http"

	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

type syncProductsHandler struct {
	productSynchronizerService *service.ProductSynchronizerService
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

	err := h.productSynchronizerService.SyncBatch(&service.SyncBatchProductsParam{
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

func newSyncProductsHandler(productSynchronizerService *service.ProductSynchronizerService) *syncProductsHandler {
	o := new(syncProductsHandler)
	o.productSynchronizerService = productSynchronizerService
	return o
}
