package http

import (
	"net/http"

	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/gin-gonic/gin"
)

type deleteProductHandler struct {
	client *typesense.Client
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

	product, err := h.client.FindProduct(req.ProductID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	product, err = h.client.DeleteProduct(product)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func newdeleteProductHandler(client *typesense.Client) *deleteProductHandler {
	o := new(deleteProductHandler)
	o.client = client
	return o
}
