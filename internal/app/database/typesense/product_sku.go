package typesense

import (
	"encoding/json"
	"github.com/AdiPP/go-typesense-product-service/internal/app/entity"
	"log"
)

func (c *Repository) FindProductSku(productSkuID int64) (result entity.ProductSku, err error) {
	// result = entity.ProductSku{}

	// if productSkuID == 0 {
	// 	err = fmt.Errorf("unknown product")
	// 	return
	// }

	// retrieveResponse, err := c.client.Collection(productSkuCollectionName.string()).Document(strconv.FormatInt(productSkuID, 10)).Retrieve(context.Background())
	// if err != nil {
	// 	return
	// }

	// fmt.Println("product document retrieved.")

	// document := newProductDocumentFromResponse(retrieveResponse)
	// result = document.transform()
	return
}

type UpsertProductSkuRequest struct {
	Document ProductSkuDocument
}

type UpsertProductSkuResponse struct {
	Document ProductSkuDocument
}

func (c *Repository) UpsertProductSku(req *UpsertProductSkuRequest) (resp UpsertProductSkuResponse, err error) {
	marshal, _ := json.Marshal(req)
	log.Printf("upserting product sku %v", string(marshal))
	return
}

func (c *Repository) DeleteProductSku(productSku entity.ProductSku) (result entity.ProductSku, err error) {
	result = entity.ProductSku{}
	return
}
