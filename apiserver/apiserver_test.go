package apiserver

import (
	"books-rest-api/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Open() error {
	return nil
}

func (m *MockStorage) Close() {

}

func TestAPIServer_Handlers(t *testing.T) {
	cfg := config.New()
	err := cfg.GetFromFile("../configs/apiserver.toml")
	if err != nil {
		t.Error(err)
	}
	s := New(cfg.Server, logrus.New(), &MockStorage{})
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
	s := New(cfg.Server, logrus.New(), &MockStorage{})

	serverDone := make(chan struct{})

	go func() {

		err := s.Run()
		if err != nil {
			// Ошибка при запуске
			t.Error(err)
			close(s.Running)
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
	s := New(cfg.Server, logrus.New(), &MockStorage{})

	s.httpServer.Addr = ":888888"

	err = s.Run()
	assert.NotNil(t, err)
}
