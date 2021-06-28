package apiserver

import (
	"books-rest-api/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAPIServer_GetAllBooks(t *testing.T) {
	c := config.New()
	err := c.EvalFromFile("../configs/apiserver.toml")
	if err != nil {
		t.Error(err)
	}
	s := New(c.Server, logrus.New())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	s.Routes()
	s.mux.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), `[{"id":1,"title":"Оно","author":"Стивен Кинг","isbn":"9-324-543-45-4"},{"id":2,"title":"Сияние","author":"Стивен Кинг","isbn":"9-326-345-66-4"}]`)
}
