package main

import (
	"log"

	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	http2 "github.com/AdiPP/go-typesense-product-service/internal/app/http"
)

func run() (err error) {
	log.Println("Initializing App...")

	typesenseClient, err := typesense.NewClient()
	if err != nil {
		return
	}

	err = http2.NewServer(typesenseClient).ListenAndServe()
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
