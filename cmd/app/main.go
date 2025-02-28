package main

import (
	"log"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	http2 "github.com/AdiPP/go-typesense-product-service/internal/app/http"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
)

func run() (err error) {
	log.Println("Initializing App...")

	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	typesenseClient, err := typesense.NewClient(cfg.Typesense)
	if err != nil {
		return
	}

	service := service.NewProductService(typesenseClient)
	server := http2.NewServer(cfg.App, typesenseClient, service)

	err = server.ListenAndServe()
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
