package server

import (
	"log"
	"net/http"
	"todo/internal/config"
	"todo/internal/intreface"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct{}

func (s *Server) DotEnvLoad() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot load .env file: %s", err)
	}
}

func setServerMode(cfg *config.Config) {
	mode := gin.ReleaseMode

	switch cfg.Project.Mode {
	case "local":
		mode = gin.DebugMode
	case "dev":
		mode = gin.TestMode
	}

	gin.SetMode(mode)
}

func runServer(cfg *config.Config) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Fatalf("Cannot run server: %s", err)
	}
}

func (s *Server) Run() {
	cfg := config.GetConfig()

	setServerMode(cfg)
	runServer(cfg)
}

func GetServer() intreface.IServer {
	return &Server{}
}
