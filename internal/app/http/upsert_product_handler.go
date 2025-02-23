package http

import (
	"net/http"

	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

type upsertProductHandler struct {
	productService *service.ProductService
}

type upsertProductRequest struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

func (h *upsertProductHandler) handle(c *gin.Context) {
	var req upsertProductRequest
	if c.ShouldBindBodyWithJSON(&req) != nil {
		c.JSON(http.StatusUnprocessableEntity, "Failed")
		return
	}

	product, err := h.productService.Upsert(&service.UpsertProductParam{
		ProductID:   req.ProductID,
		ProductName: req.ProductName,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func newUpsertProductHandler(productService *service.ProductService) *upsertProductHandler {
	o := new(upsertProductHandler)
	o.productService = productService
	return o
}
