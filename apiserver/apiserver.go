package apiserver

import (
	"books-rest-api/config"
	"books-rest-api/storage"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*// IStorage is an interface to the storage
type IStorage interface {
	Add(models.Book)
	GetAll() models.Books
}*/

// Server ...
type Server struct {
	httpServer *http.Server
	config     *config.ServerConfig
	logger     *logrus.Logger
	mux        *gin.Engine
	storage    *storage.Storage
	Running    chan struct{}
}

// New ...
func New(c *config.ServerConfig, l *logrus.Logger, stor *storage.Storage) *Server {

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
