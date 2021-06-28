package apiserver

// Routes ...
func (s *Server) Routes() {
	s.mux.GET("/books", s.GetAll())
	s.mux.GET("/books/:id", s.Get())
	s.mux.POST("/books", s.AddNew())
	s.mux.PUT("/books/:id", s.Replace())
	s.mux.PATCH("/books/:id", s.Update())
	s.mux.DELETE("/books/:id", s.Delete())
}
