package server

import (
	"log"
	"todo/internal/common"

	"github.com/joho/godotenv"
)

func envLoad() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func Run() {
	envLoad()
	logger := common.GetLogger()
	logger.Debug("Hello, World!")
}
