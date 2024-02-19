package server

import (
	"fmt"
	"log"
	"todo/internal/config"

	"github.com/joho/godotenv"
)

func envLoad() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func Run() {
	envLoad()
	cfg := config.GetConfig()
	fmt.Print(cfg)
}
