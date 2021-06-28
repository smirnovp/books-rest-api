package apiserver

import (
	"books-rest-api/config"
	"books-rest-api/storage"
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
	config  *config.ServerConfig
	logger  *logrus.Logger
	mux     *gin.Engine
	storage *storage.Storage
	Running chan struct{}
}

// New ...
func New(c *config.ServerConfig, l *logrus.Logger) *Server {
	return &Server{
		config:  c,
		logger:  l,
		mux:     gin.Default(),
		Running: make(chan struct{}),
	}
}

// Run ...
func (s *Server) Run() error {

	server := &http.Server{
		Addr:           s.config.Addr,
		Handler:        s.mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.Routes()

	s.logger.Infof("Server is running on port %s ...", s.config.Addr)

	close(s.Running)
	return server.ListenAndServe()
}
