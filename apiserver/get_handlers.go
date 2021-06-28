package apiserver

import (
	"books-rest-api/models"

	"github.com/gin-gonic/gin"
)

// GetAll gets all books from the storage
func (s *Server) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, models.Books{
			models.Book{
				ID:     1,
				Title:  "Оно",
				Author: "Стивен Кинг",
				ISBN:   "9-324-543-45-4",
			},
			models.Book{
				ID:     2,
				Title:  "Сияние",
				Author: "Стивен Кинг",
				ISBN:   "9-326-345-66-4",
			},
		})
	}
}

// Get ...
func (s *Server) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
