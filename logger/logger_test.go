package logger

import (
	"books-rest-api/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	_, err := New(&config.LoggerConfig{
		Level: "debug",
	})
	assert.Nil(t, err)

	_, err = New(&config.LoggerConfig{
		Level: "badlevel",
	})
	assert.NotEqual(t, nil, err)
}
