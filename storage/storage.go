package storage

import "database/sql"

// Storage ...
type Storage struct {
	config *Config
	db     *sql.DB
}

// New ...
func New(c *Config) *Storage {
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
