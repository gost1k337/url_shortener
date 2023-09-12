package main

import (
	"github.com/gost1k337/url_shortener/user_service/config"
	"github.com/gost1k337/url_shortener/user_service/internal/app"
	"log"
)

const configPath = "config/config.yaml"

func main() {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
