package main

import (
	"books-rest-api/apiserver"
	"books-rest-api/config"
	"books-rest-api/logger"
	"flag"
	"log"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "configs/apiserver.toml", "путь к файлу конфигурации API-сервера")
	flag.Parse()

	config := config.New()
	if err := config.EvalFromFile(configFile); err != nil {
		log.Fatal("Could not populate the config structure from a config file: ", err)
	}

	logger, err := logger.NewLogger(config.Logger)
	if err != nil {
		log.Fatal("Could not get logger: ", err)
	}

	apiserver := apiserver.New(config.Server, logger)
	err = apiserver.Run()
	if err != nil {
		log.Fatal(err)
	}
}
