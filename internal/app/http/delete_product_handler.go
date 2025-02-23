package http

import (
	"net/http"

	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

type deleteProductHandler struct {
	productService *service.ProductService
}

type deleteProductRequest struct {
	ProductID int64 `json:"product_id"`
}

func (h *deleteProductHandler) handle(c *gin.Context) {
	var req deleteProductRequest
	if c.ShouldBindBodyWithJSON(&req) != nil {
		c.String(http.StatusUnprocessableEntity, "Failed")
		return
	}

	product, err := h.productService.Delete(&service.DeleteProductParam{
		ProductID: req.ProductID,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func newdeleteProductHandler(productService *service.ProductService) *deleteProductHandler {
	o := new(deleteProductHandler)
	o.productService = productService
	return o
}
