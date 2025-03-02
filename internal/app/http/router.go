package http

func (s *Server) initRouter() {
	s.engine.GET("/ping", newPingHandler().handle)

	s.engine.POST("/products/sync", newSyncProductsHandler(s.psService).handle)
}
