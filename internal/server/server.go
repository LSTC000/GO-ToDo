package server

import (
	"fmt"
	_ "todo/docs"
	todoV1 "todo/internal/api/todo/v1"
	"todo/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IServer interface {
	Run() error
}

type Server struct{}

func setDotEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("cannot load .env file: %w", err)
	}
	return nil
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

func setSwagger(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setV1Handlers(r *gin.Engine) {
	rg := r.Group("/api/v1")

	todoV1Handler := todoV1.GetHandler()
	todoV1Handler.Register(rg)
}

func runServer(r *gin.Engine, cfg *config.Config) error {
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		return fmt.Errorf("cannot run main router: %w", err)
	}
	return nil
}

func (s *Server) Run() error {
	r := gin.Default()
	cfg := config.GetConfig()

	if err := setDotEnv(); err != nil {
		return err
	}

	setServerMode(cfg)
	setSwagger(r)
	setV1Handlers(r)

	if err := runServer(r, cfg); err != nil {
		return err
	}

	return nil
}

func GetServer() IServer {
	return &Server{}
}
