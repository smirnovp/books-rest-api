package apiserver

import (
	"books-rest-api/config"
	"books-rest-api/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// IStorage is an interface to the storage
type IStorage interface {
	Open() error
	Close() error
	Add(models.Book) (int64, error)
	GetAll() (models.Books, error)
	Delete(int64) (int64, error)
	Update(int64, models.Book) (int64, error)
}

// Server ...
type Server struct {
	httpServer *http.Server
	config     *config.ServerConfig
	logger     *logrus.Logger
	mux        *gin.Engine
	storage    IStorage
	Running    chan struct{}
}

// New ...
func New(c *config.ServerConfig, l *logrus.Logger, stor IStorage) *Server {

	gin.SetMode(c.GinMode)
	mux := gin.Default()

	return &Server{
		httpServer: &http.Server{
			Addr:           c.Addr,
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		config:  c,
		logger:  l,
		mux:     mux,
		storage: stor,
		Running: make(chan struct{}),
	}
}

// Run ...
func (s *Server) Run() error {

	err := s.storage.Open()
	if err != nil {
		return err
	}

	s.Routes()

	s.logger.Infof("Server is running on port %s ...", s.config.Addr)

	close(s.Running)
	err = s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop ...
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
