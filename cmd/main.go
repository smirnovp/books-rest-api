package main

import (
	"books-rest-api/apiserver"
	"books-rest-api/config"
	"books-rest-api/logger"
	"books-rest-api/storage"
	"flag"
	"log"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "configs/apiserver.toml", "путь к файлу конфигурации API-сервера")
	flag.Parse()

	cfg := config.New()
	if err := cfg.GetFromFile(configFile); err != nil {
		log.Fatal("Could not get config from the file: ", err)
	}

	logger, err := logger.New(cfg.Logger)
	if err != nil {
		log.Fatal("Could not get logger: ", err)
	}

	storage := storage.New(cfg.Storage)

	apiserver := apiserver.New(cfg.Server, logger, storage)
	err = apiserver.Run()
	if err != nil {
		log.Fatal(err)
	}
}
