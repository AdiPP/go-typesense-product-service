package typesense

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AdiPP/go-typesense-product-service/internal/app"
	"github.com/typesense/typesense-go/v3/typesense/api"
	"github.com/typesense/typesense-go/v3/typesense/api/pointer"
)

type productDocument struct {
	ID          string `json:"id"`
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

func (d *productDocument) transform() app.Product {
	return app.Product{
		ProductID:   d.ProductID,
		ProductName: d.ProductName,
	}
}

func newProductDocument(product app.Product) productDocument {
	return productDocument{
		ID:          strconv.FormatInt(product.ProductID, 10),
		ProductID:   product.ProductID,
		ProductName: product.ProductName,
	}
}

func (c *Client) initProductSchema() (err error) {
	collection, err := c.client.Collection(productCollectionName.string()).Retrieve(context.Background())
	if err != nil && collection != nil {
		return
	}

	if collection != nil {
		return
	}

	schema := &api.CollectionSchema{
		Name: productCollectionName.string(),
		Fields: []api.Field{
			{Name: "product_id", Type: "int64"},
			{Name: "product_name", Type: "string"},
		},
		DefaultSortingField: pointer.String("product_id"),
	}

	_, err = c.client.Collections().Create(context.Background(), schema)
	if err != nil {
		return
	}

	fmt.Println("product schema created.")
	return
}

func (c *Client) UpsertProduct(product app.Product) (result app.Product, err error) {
	result = app.Product{}

	if product.ProductID == 0 {
		err = fmt.Errorf("unknown product")
		return
	}

	upsertResult, err := c.client.Collection("products").Documents().Upsert(context.Background(), newProductDocument(product), &api.DocumentIndexParameters{})
	if err != nil {
		return
	}

	fmt.Println("product document upserted.")

	document := productDocument{
		ID:          upsertResult["id"].(string),
		ProductID:   int64(upsertResult["product_id"].(float64)),
		ProductName: upsertResult["product_name"].(string),
	}
	result = document.transform()
	return
}
