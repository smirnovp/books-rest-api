package logger

import (
	"books-rest-api/config"

	"github.com/sirupsen/logrus"
)

// New ...
func New(c *config.LoggerConfig) (*logrus.Logger, error) {
	l := logrus.New()
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}
	l.SetLevel(level)
	return l, nil
}
