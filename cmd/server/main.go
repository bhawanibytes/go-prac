package main

import (
	"go-prac/internal/logger"
	"go-prac/internal/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger.Init()
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	server.Start(port)
}
