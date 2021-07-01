package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAll gets all books from the storage
func (s *Server) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		bs, err := s.storage.GetAll()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %s", err.Error())
			return
		}

		c.JSON(200, bs)
	}
}

// Get ...
func (s *Server) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
