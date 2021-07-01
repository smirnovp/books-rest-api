package apiserver

import (
	"books-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Replace ...
func (s *Server) Replace() gin.HandlerFunc {
	type URLparam struct {
		ID int64 `uri:"id" binding:"required"`
	}
	return func(c *gin.Context) {
		var prm URLparam
		err := c.ShouldBindUri(&prm)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var b models.Book
		err = c.ShouldBindJSON(&b)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		affectedRows, err := s.storage.Update(prm.ID, b)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if affectedRows < 1 {
			c.Status(http.StatusConflict)
			return
		}
		c.Status(http.StatusNoContent)
	}
}
