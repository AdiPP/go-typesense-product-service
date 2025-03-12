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

func NewRepository(cfg *config.Config) (c *Repository, err error) {
	c = &Repository{
		client: typesense.NewClient(
			typesense.WithNodes([]string{
				fmt.Sprintf("%s:%s", cfg.TypesenseDatabaseHost, cfg.TypesenseDatabasePort),
			}),
			typesense.WithAPIKey(cfg.TypesenseDatabaseAPIKey),
			typesense.WithConnectionTimeout(2*time.Second),
		),
	}

	if ok, err := c.client.Health(context.Background(), 2*time.Second); err != nil || !ok {
		log.Fatal(fmt.Errorf("could not connect to typesense database: %w", err))
	}

	log.Print("Success connect to typesense client")
	return
}
