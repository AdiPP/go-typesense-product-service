package typesense

import (
	"context"

	"github.com/typesense/typesense-go/v3/typesense/api"
	"github.com/typesense/typesense-go/v3/typesense/api/pointer"
)

type productDocument struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

func (c *TypesenseClient) initProductSchema() {
	schema := &api.CollectionSchema{
		Name: "products",
		Fields: []api.Field{
			{Name: "product_id", Type: "int64"},
			{Name: "product_name", Type: "string"},
		},
		DefaultSortingField: pointer.String("product_id"),
	}
	typesenseClient := c.typesenseClient

	typesenseClient.Collections().Create(context.Background(), schema)
}
