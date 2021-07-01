package apiserver

import (
	"books-rest-api/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAPIServer_Handlers(t *testing.T) {
	cfg := config.New()
	err := cfg.GetFromFile("../configs/apiserver.toml")
	if err != nil {
		t.Error(err)
	}
	s := New(cfg.Server, logrus.New())
	s.Routes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	s.mux.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), `[{"id":1,"title":"Оно","author":"Стивен Кинг","isbn":"9-324-543-45-4"},{"id":2,"title":"Сияние","author":"Стивен Кинг","isbn":"9-326-345-66-4"}]`)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/books", nil)
	s.mux.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 204)

}

func TestAPIServerRUN_OK(t *testing.T) {
	cfg := config.New()
	err := cfg.GetFromFile("../configs/apiserver.toml")
	if err != nil {
		t.Error(err)
	}
	s := New(cfg.Server, logrus.New())

	serverDone := make(chan struct{})

	go func() {

		err := s.Run()
		if err != nil {
			// Ошибка при запуске
			t.Error(err)
		}
		// В этой точке сервер остановлен err == nil
		defer close(serverDone)

	}()

	<-s.Running // Дожидаемся запуска сервера (канал закроется, перед самым запуском сервера)

	err = s.Stop()
	if err != nil {
		t.Error("Ошибка остановки сервера: ", err)
	}

	<-serverDone // Обязательно дожидаемся полной остановки сервера
}

func TestAPIServerRUN_Fail(t *testing.T) {
	cfg := config.New()
	err := cfg.GetFromFile("../configs/apiserver.toml")
	if err != nil {
		t.Error(err)
	}
	s := New(cfg.Server, logrus.New())

	s.httpServer.Addr = ":888888"

	err = s.Run()
	assert.NotNil(t, err)
}
