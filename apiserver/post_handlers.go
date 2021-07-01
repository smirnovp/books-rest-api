package apiserver

import (
	"books-rest-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddNew adds new book in the storage
func (s *Server) AddNew() gin.HandlerFunc {
	var b models.Book
	return func(c *gin.Context) {
		err := c.ShouldBind(&b)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := s.storage.Add(b)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		b.ID = id

		c.Header("Location", fmt.Sprintf("/books/%d", id))

		c.JSON(http.StatusCreated, b)
	}
}
