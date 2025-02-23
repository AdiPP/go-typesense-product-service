package http

import (
	"net/http"

	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/AdiPP/go-typesense-product-service/internal/app/entity"
	"github.com/gin-gonic/gin"
)

type upsertProductHandler struct {
	client *typesense.Client
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

	product, err := h.client.UpsertProduct(entity.Product{
		ProductID:   req.ProductID,
		ProductName: req.ProductName,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func newUpsertProductHandler(client *typesense.Client) *upsertProductHandler {
	o := new(upsertProductHandler)
	o.client = client
	return o
}
