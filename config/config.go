package config

import (
	"github.com/BurntSushi/toml"
)

// LoggerConfig ...
type LoggerConfig struct {
	Level string `toml:"level"`
}

// ServerConfig ...
type ServerConfig struct {
	Addr string `toml:"addr"`
}

// StorageConfig ...
type StorageConfig struct {
	DatabaseURL string `toml:"database_url"`
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
		Logger: &LoggerConfig{
			Level: "debug",
		},
		Server: &ServerConfig{
			Addr: ":8082",
		},
		Storage: &StorageConfig{
			DatabaseURL: "some default URL string",
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
