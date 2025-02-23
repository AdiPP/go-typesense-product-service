package http

import (
	"net/http"

	"github.com/AdiPP/go-typesense-product-service/internal/app"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/gin-gonic/gin"
)

type addProductHandler struct {
	client *typesense.Client
}

type addProductRequest struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

func (h *addProductHandler) handle(c *gin.Context) {
	var req addProductRequest
	if c.ShouldBindBodyWithJSON(&req) != nil {
		c.String(http.StatusUnprocessableEntity, "Failed")
		return
	}

	product, err := h.client.UpsertProduct(app.Product{
		ProductID:   req.ProductID,
		ProductName: req.ProductName,
	})
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func newAddProductHandler(client *typesense.Client) *addProductHandler {
	o := new(addProductHandler)
	o.client = client
	return o
}
