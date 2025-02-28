package main

import (
	"fmt"
	"log"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/pgsql"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/AdiPP/go-typesense-product-service/internal/app/http"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
)

func run() (err error) {
	log.Println("Initializing App...")

	cfg, err := config.LoadConfig()
	if err != nil {
		err = fmt.Errorf("failed to load config: %s", err)
		return
	}

	typesenseRepo, err := typesense.NewClient(cfg)
	if err != nil {
		err = fmt.Errorf("failed to load typesense repo: %s", err)
		return
	}

	pgsqlRepo, err := pgsql.NewRepository(cfg)
	if err != nil {
		err = fmt.Errorf("failed to load pgsql repo: %s", err)
		return
	}

	productService := service.NewProductService(typesenseRepo)
	productSynchorizerService := service.NewProductSynchorizerService(pgsqlRepo)

	err = http.NewServer(cfg, pgsqlRepo, typesenseRepo, productService, productSynchorizerService).ListenAndServe()
	if err != nil {
		return
	}
	return
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
