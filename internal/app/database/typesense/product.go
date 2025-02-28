package typesense

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AdiPP/go-typesense-product-service/internal/app/entity"
	"github.com/typesense/typesense-go/v3/typesense/api"
	"github.com/typesense/typesense-go/v3/typesense/api/pointer"
)

type productDocument struct {
	ID          string `json:"id"`
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
}

func (d *productDocument) productIdString() string {
	return strconv.FormatInt(d.ProductID, 10)
}

func (d *productDocument) transform() entity.Product {
	return entity.Product{
		ProductID:   d.ProductID,
		ProductName: d.ProductName,
	}
}

func newProductDocument(product entity.Product) productDocument {
	return productDocument{
		ID:          strconv.FormatInt(product.ProductID, 10),
		ProductID:   product.ProductID,
		ProductName: product.ProductName,
	}
}

func newProductDocumentFromResponse(response map[string]any) productDocument {
	return productDocument{
		ID:          response["id"].(string),
		ProductID:   int64(response["product_id"].(float64)),
		ProductName: response["product_name"].(string),
	}
}

func (c *Repository) initProductSchema() (err error) {
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

func (c *Repository) FindProduct(productID int64) (result entity.Product, err error) {
	result = entity.Product{}

	if productID == 0 {
		err = fmt.Errorf("unknown product")
		return
	}

	retrieveResponse, err := c.client.Collection(productCollectionName.string()).Document(strconv.FormatInt(productID, 10)).Retrieve(context.Background())
	if err != nil {
		return
	}

	fmt.Println("product document retrieved.")

	document := newProductDocumentFromResponse(retrieveResponse)
	result = document.transform()
	return
}

func (c *Repository) UpsertProduct(product entity.Product) (result entity.Product, err error) {
	result = entity.Product{}

	if product.ProductID == 0 {
		err = fmt.Errorf("unknown product")
		return
	}

	document := newProductDocument(product)
	upsertResponse, err := c.client.Collection(productCollectionName.string()).Documents().Upsert(context.Background(), document, &api.DocumentIndexParameters{})
	if err != nil {
		return
	}

	fmt.Println("product document upserted.")

	document = newProductDocumentFromResponse(upsertResponse)
	result = document.transform()
	return
}

func (c *Repository) DeleteProduct(product entity.Product) (result entity.Product, err error) {
	result = entity.Product{}

	if product.ProductID == 0 {
		err = fmt.Errorf("unknown product")
		return
	}

	document := newProductDocument(product)
	deleteResponse, err := c.client.Collection(productCollectionName.string()).Document(document.productIdString()).Delete(context.Background())
	if err != nil {
		return
	}

	fmt.Println("product document deleted.")

	document = newProductDocumentFromResponse(deleteResponse)
	result = document.transform()
	return
}
