package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine          *gin.Engine
	typesenseClient *typesense.Client
}

func (r *Server) initRouter() {
	r.engine.GET("/ping", newPingHandler().handle)
	r.engine.POST("/products", newAddProductHandler(r.typesenseClient).handle)
}

func (s *Server) ListenAndServe() (err error) {
	port := "8000"

	log.Printf("App running on port : %v \n", port)

	server := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%v", port),
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

func NewServer(typesenseClient *typesense.Client) *Server {
	o := new(Server)
	o.engine = gin.Default()
	o.typesenseClient = typesenseClient

	o.initRouter()

	return o
}
