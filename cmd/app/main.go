package main

import (
	"log"

	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	http2 "github.com/AdiPP/go-typesense-product-service/internal/app/http"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
)

func run() (err error) {
	log.Println("Initializing App...")

	typesenseClient, err := typesense.NewClient()
	if err != nil {
		return
	}

	service := service.NewService(typesenseClient)
	
	err = http2.NewServer(typesenseClient, service).ListenAndServe()
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
