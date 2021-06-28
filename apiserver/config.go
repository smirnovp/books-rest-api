package apiserver

import (
	"books-rest-api/storage"
	"log"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	Addr     string          `toml:"addr"`
	LogLevel string          `toml:"log_level"`
	Storage  *storage.Config `toml:"storage"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Addr:     ":8080",
		LogLevel: "debug",
		Storage:  storage.NewConfig(),
	}
}

// EvalFromFile ...
func (c *Config) EvalFromFile(filename string) {
	config := Config{}
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		log.Println("Ошибка файла конфигурации: ", err)
		log.Println("Использую дефолтные значения")
		return
	}
	*c = config
}
