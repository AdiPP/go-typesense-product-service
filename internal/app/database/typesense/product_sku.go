package typesense

import (
	"context"
	"encoding/json"
	"github.com/typesense/typesense-go/v3/typesense/api"
)

type UpsertProductSkuRequest struct {
	Document ProductSkuDocument
}

type UpsertProductSkuResponse struct {
	Document ProductSkuDocument
}

func (c *Repository) UpsertProductSku(req *UpsertProductSkuRequest) (resp UpsertProductSkuResponse, err error) {
	resp = UpsertProductSkuResponse{
		Document: ProductSkuDocument{},
	}

	if nil == req {
		return
	}

	schema := newSchemaGetter().getProductSkuSchema()
	upsertResponse, err := c.client.Collection(schema.Name).Documents().Upsert(
		context.Background(),
		req.Document,
		&api.DocumentIndexParameters{},
	)

	if err != nil {
		return
	}

	jsonData, err := json.Marshal(upsertResponse)
	if err != nil {
		return
	}

	var document ProductSkuDocument
	err = json.Unmarshal(jsonData, &document)
	if err != nil {
		return
	}

	resp = UpsertProductSkuResponse{
		Document: document,
	}
	return
}
