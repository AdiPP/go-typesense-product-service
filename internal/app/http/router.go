package http

func (r *Server) initRouter() {
	r.engine.GET("/ping", newPingHandler().handle)
	r.engine.PATCH("/products", newUpsertProductHandler(r.productService).handle)
	r.engine.DELETE("/products", newdeleteProductHandler(r.productService).handle)

	r.engine.POST("/products/sync", newSyncProductHandler(r.productSynchorizerService).handle)
}
