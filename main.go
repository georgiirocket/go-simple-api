package main

import (
	"github.com/joho/godotenv"
	"go-simple/docs"
	"log/slog"
	"os"
)

func main() {
	loadEnv()

	docs.SwaggerInfo.Title = "Go simple api"
	docs.SwaggerInfo.Version = "1.0"
}

func loadEnv() {
	err := godotenv.Load(".env." + os.Getenv("MODE"))

	if err != nil {
		panic("Error loading env file")
	}

	slog.Info("Loaded env successfully")
}
