package http

import (
	"net/http"

	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

type syncProductHandler struct {
	productSynchorizerService *service.ProductSynchorizerService
}

type syncProductRequest struct {
	ProductIDs []int64 `json:"product_ids"`
}

func (h *syncProductHandler) handle(c *gin.Context) {
	var req syncProductRequest
	if c.ShouldBindBodyWithJSON(&req) != nil {
		c.String(http.StatusUnprocessableEntity, "Failed")
		return
	}

	err := h.productSynchorizerService.SyncBatch(&service.SyncProductBatchParam{
		ProductIDs: req.ProductIDs,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func newSyncProductHandler(productSynchorizerService *service.ProductSynchorizerService) *syncProductHandler {
	o := new(syncProductHandler)
	o.productSynchorizerService = productSynchorizerService
	return o
}
