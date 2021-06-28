package storage

import (
	"books-rest-api/config"
	"database/sql"
)

// Storage ...
type Storage struct {
	config *config.StorageConfig
	db     *sql.DB
}

// New ...
func New(c *config.StorageConfig) *Storage {
	return &Storage{
		config: c,
	}
}

// Open ...
func (stor *Storage) Open() error {
	return nil
}

// Close ...
func (stor *Storage) Close() {
	//...
}
