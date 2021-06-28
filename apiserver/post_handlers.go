package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddNew adds new book in the storage
func (s *Server) AddNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	}
}
