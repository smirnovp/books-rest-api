package config

import (
	"github.com/BurntSushi/toml"
)

// ServerConfig ...
type ServerConfig struct {
	Addr    string `toml:"addr"`
	GinMode string `toml:"gin_mode"`
}

// LoggerConfig ...
type LoggerConfig struct {
	Level string `toml:"level"`
}

// StorageConfig ...
type StorageConfig struct {
	DatabaseURL string `toml:"database_url"`
}

// GinConfig ...
type GinConfig struct {
	Mode string `toml:"mode"`
}

// Config ...
type Config struct {
	Server  *ServerConfig  `toml:"server"`
	Logger  *LoggerConfig  `toml:"logger"`
	Storage *StorageConfig `toml:"storage"`
}

// New ...
func New() *Config {
	return &Config{
		Server: &ServerConfig{
			Addr:    ":8082",
			GinMode: "release",
		},
		Logger: &LoggerConfig{
			Level: "debug",
		},
		Storage: &StorageConfig{
			DatabaseURL: "host=localhost dbname=books-rest-api sslmode=disable",
		},
	}
}

// GetFromFile ...
func (c *Config) GetFromFile(filename string) error {
	_, err := toml.DecodeFile(filename, &c)
	if err != nil {
		return err
	}
	return nil
}
