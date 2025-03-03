package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg             config.App
	engine          *gin.Engine
	typesenseClient *typesense.Client
	productService  *service.ProductService
}

func (s *Server) ListenAndServe() (err error) {
	host := s.cfg.Host
	port := s.cfg.Port

	log.Printf("App running on %s:%v \n", host, port)

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%v", host, port),
		Handler:           s.engine.Handler(),
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		err = fmt.Errorf("failed running App on port : %v. error: %v", port, err)
		return
	}

	return
}

func NewServer(cfg config.App, typesenseClient *typesense.Client, productService *service.ProductService) *Server {
	switch cfg.GetEnv() {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "staging":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	s := new(Server)
	s.cfg = cfg
	s.engine = gin.Default()
	s.typesenseClient = typesenseClient
	s.productService = productService

	s.initRouter()

	return s
}
