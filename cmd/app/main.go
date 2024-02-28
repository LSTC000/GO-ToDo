package main

import (
	"log"
	"todo/internal/server"
)

// @title           ToDo
// @version         1.0.0
// @description     This is a ToDo API.

// @host      0.0.0.0:8000
// @BasePath  /api/v1

func main() {
	s := server.GetServer()
	if err := s.Run(); err != nil {
		log.Fatalf("Run server error: %v", err)
	}
}
