package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddBook adds new book in the storage
func (s *Server) AddBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	}
}
