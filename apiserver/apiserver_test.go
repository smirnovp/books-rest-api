package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_GetAllBooks(t *testing.T) {
	s := New(NewConfig())
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	s.Routes()
	s.mux.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), `[{"id":1,"title":"Оно","author":"Стивен Кинг","isbn":"9-324-543-45-4"},{"id":2,"title":"Сияние","author":"Стивен Кинг","isbn":"9-326-345-66-4"}]`)
}
