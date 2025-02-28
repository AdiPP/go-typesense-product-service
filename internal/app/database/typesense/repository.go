package typesense

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/typesense/typesense-go/v3/typesense"
)

type Repository struct {
	client *typesense.Client
}

func (c *Repository) init() (err error) {
	ok, err := c.client.Health(context.Background(), 2*time.Second)
	if err != nil {
		return
	}

	if !ok {
		err = fmt.Errorf("failed connect to typesense client")
		return
	}

	log.Println("Success connect to typesense client")

	// err = c.initProductSchema()
	return
}

func NewClient(cfg *config.Config) (c *Repository, err error) {
	c = &Repository{
		client: typesense.NewClient(
			typesense.WithNodes([]string{
				fmt.Sprintf("%s:%s", cfg.TypesenseDatabaseHost, cfg.TypesenseDatabasePort),
			}),
			typesense.WithAPIKey(cfg.TypesenseDatabaseAPIKey),
			typesense.WithConnectionTimeout(2*time.Second),
		),
	}

	err = c.init()
	return
}
