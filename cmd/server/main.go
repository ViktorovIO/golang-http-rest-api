package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"http-rest-api/internal/app/server"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	apiServer := server.New(config)
	if err := apiServer.Start(); err != nil {
		log.Fatal(err)
	}
}
