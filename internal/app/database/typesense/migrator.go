package typesense

import (
	"context"
	"log"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
)

type Migrator struct {
	repo *Repository
}

func NewMigrator(cfg *config.Config) (*Migrator, error) {
	repo, err := NewRepository(cfg)
	if err != nil {
		return nil, err
	}

	return &Migrator{repo: repo}, nil
}

func (m *Migrator) Migrate() (err error) {
	err = m.migrateProductSkuSchema()
	return
}

func (m *Migrator) migrateProductSkuSchema() (err error) {
	schema := newSchemaGetter().getProductSkuSchema()

	collection, err := m.repo.client.Collection(schema.Name).Retrieve(context.Background())
	if err != nil && collection != nil {
		return
	}

	if collection != nil {
		return
	}

	_, err = m.repo.client.Collections().Create(context.Background(), schema)
	if err != nil {
		return
	}

	log.Println("product schema created.")
	return
}
