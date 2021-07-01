package apiserver

// Routes ...
func (s *Server) Routes() {
	books := s.mux.Group("/books")
	{
		books.GET("", s.GetAll())
		books.GET("/:id", s.Get())
		books.POST("", s.AddNew())
		books.PUT("/:id", s.Replace())
		books.PATCH("/:id", s.Update())
		books.DELETE("/:id", s.Delete())
	}
	/*s.mux.GET("/books", s.GetAll())
	s.mux.GET("/books/:id", s.Get())
	s.mux.POST("/books", s.AddNew())
	s.mux.PUT("/books/:id", s.Replace())
	s.mux.PATCH("/books/:id", s.Update())
	s.mux.DELETE("/books/:id", s.Delete())*/
}
