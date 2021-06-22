package main

import (
	"books-rest-api/apiserver"
	"flag"
	"log"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "configs/apiserver.toml", "путь к файлу конфигурации API-сервера")
	flag.Parse()

	config := apiserver.NewConfig()
	config.EvalFromFile(configFile)

	server := apiserver.New(config)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
