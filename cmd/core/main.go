package main

import (
	"go-simple-api/config"
	"go-simple-api/utils/server"
	"log"
)

func main() {
	config.Init()

	app := server.NewApp()

	if err := app.Run(config.Env.AppPort); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
