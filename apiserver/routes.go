package apiserver

// Routes ...
func (s *Server) Routes() {
	s.mux.GET("/books", s.GetAllBooks())
	s.mux.POST("/books", s.AddBook())
}
