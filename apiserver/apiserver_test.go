package apiserver

import (
	"books-rest-api/config"
	"books-rest-api/models"
	"bytes"
	"encoding/json"
	"fmt"
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

func (m *MockStorage) Close() error {
	return nil
}

func (m *MockStorage) GetAll() (models.Books, error) {
	bs := models.Books{
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
	}
	return bs, nil
}

func (m *MockStorage) Add(b models.Book) (int64, error) {
	return 1, nil
}

func (m *MockStorage) Delete(id int64) (int64, error) {
	return 1, nil
}
func (m *MockStorage) Update(id int64, b models.Book) (int64, error) {
	return 1, nil
}

func TestAPIServer_Handlers(t *testing.T) {
	b := models.Book{
		Title:  "Оно",
		Author: "Стивен Кинг",
		ISBN:   "9-324-543-45-4",
	}
	bs := models.Books{
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
	}
	bsJSON, err := json.Marshal(bs)
	if err != nil {
		t.Error(err)
	}

	bJSON, err := json.Marshal(b)
	if err != nil {
		t.Error(err)
	}

	cfg := config.New()

	err = cfg.GetFromFile("../configs/apiserver.toml")
	if err != nil {
		t.Error(err)
	}

	s := New(cfg.Server, logrus.New(), &MockStorage{})

	s.Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	s.mux.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, string(bsJSON), w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/books", bytes.NewReader(bJSON))
	req.Header.Set("Content-Type", "application/json")
	s.mux.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated, fmt.Sprint(w.Body.String()))

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
