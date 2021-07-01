package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete ...
func (s *Server) Delete() gin.HandlerFunc {
	type URIparam struct {
		ID int64 `uri:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		var prm URIparam

		err := c.ShouldBindUri(&prm)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		n, err := s.storage.Delete(prm.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if n < 1 {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}
