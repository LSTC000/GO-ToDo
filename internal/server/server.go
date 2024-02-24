package server

import (
	"log"
	"todo/internal/config"
	"todo/internal/intreface"
	todo "todo/internal/todo/v1"

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

func setV1Handlers(r *gin.Engine) {
	rg := r.Group("/v1")

	todoHandler := todo.GetHandler()
	todoHandler.Register(rg)
}

func runServer(r *gin.Engine, cfg *config.Config) {
	err := r.Run(":" + cfg.Server.Port)
	if err != nil {
		log.Fatalf("Cannot run main router: %s", err)
	}
}

func (s *Server) Run() {
	r := gin.Default()
	cfg := config.GetConfig()

	setServerMode(cfg)
	setV1Handlers(r)
	runServer(r, cfg)
}

func GetServer() intreface.IServer {
	return &Server{}
}
