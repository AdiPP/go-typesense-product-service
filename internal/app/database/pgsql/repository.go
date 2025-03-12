package pgsql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

type Repository struct {
	database *goqu.Database
}

func NewRepository(cfg *config.Config) (repo *Repository, err error) {
	const dialect = "postgres"

	db, err := sql.Open(
		dialect,
		cfg.GetPgsqlConnString(),
	)
	if err != nil {
		log.Fatal(fmt.Errorf("error open to pgsql database: %w", err))
	}

	if err = db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("error connecting to pgsql database: %w", err))
	}

	log.Print("success connect to pgsql database")

	repo = &Repository{
		database: goqu.Dialect(dialect).DB(db),
	}
	return
}
