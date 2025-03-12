package main

import (
	"github.com/AdiPP/go-typesense-product-service/internal/app/service/cdc"
	"log"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/pgsql"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/AdiPP/go-typesense-product-service/internal/app/http"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
)

func run() (err error) {
	log.Print("Starting APP...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	tpRepo, err := typesense.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed to load typesense repo: %s", err)
	}

	pgsqlRepo, err := pgsql.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed to load pgsql repo: %s", err)
	}

	psService := service.NewProductSynchronizerService(pgsqlRepo, tpRepo)
	cdcService := cdc.NewService(cfg, psService)

	go cdcService.StartCDC()

	if err = http.NewServer(cfg, pgsqlRepo, tpRepo, psService).ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
	return
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
