package storage

import (
	"books-rest-api/config"
	"database/sql"

	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", stor.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	stor.db = db

	return nil
}

// Close ...
func (stor *Storage) Close() {
	stor.db.Close()
}
