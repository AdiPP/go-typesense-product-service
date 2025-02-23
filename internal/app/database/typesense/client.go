package typesense

import (
	"time"

	"github.com/typesense/typesense-go/v3/typesense"
)

type TypesenseClient struct {
	typesenseClient *typesense.Client
}

func NewTypesenseClient() *TypesenseClient {
	client := &TypesenseClient{
		typesenseClient: typesense.NewClient(
			typesense.WithNodes([]string{
				"http://localhost:8108",
			}),
			typesense.WithAPIKey("xyz"),
			typesense.WithConnectionTimeout(2*time.Second),
		),
	}

	client.initProductSchema()

	return client
}
