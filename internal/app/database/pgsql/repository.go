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

func (r *Repository) init() (err error) {

	return
}

func NewRepository(cfg *config.Config) (repo *Repository, err error) {
	dialect := goqu.Dialect("postgres")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		cfg.PgsqlDatabaseHost,
		cfg.PgsqlDatabasePort,
		cfg.PgsqlDatabaseUsername,
		cfg.PgsqlDatabasePassword,
		cfg.PgsqlDatabaseDatabase,
		cfg.PgsqlDatabaseSchema,
	)

	pgDb, err := sql.Open(
		"postgres",
		dsn,
	)
	if err != nil {
		return
	}

	err = pgDb.Ping()
	if err != nil {
		return
	}

	log.Println("Success connect to pgsql database")

	repo = &Repository{
		database: dialect.DB(pgDb),
	}
	return
}
