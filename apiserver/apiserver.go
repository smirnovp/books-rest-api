package apiserver

import (
	"books-rest-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	config  *Config
	logger  *logrus.Logger
	mux     *gin.Engine
	storage models.IStorage
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		mux:    gin.Default(),
	}
}

// Run ...
func (s *Server) Run() error {
	if err := s.configLogLevel(); err != nil {
		return err
	}

	server := &http.Server{
		Addr:           s.config.Addr,
		Handler:        s.mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.Routes()
	s.logger.Debugf("Server is running on port %s ...", s.config.Addr)
	return server.ListenAndServe()

}

func (s *Server) configLogLevel() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
