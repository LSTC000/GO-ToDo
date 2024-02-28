package v1

import (
	"todo/internal/api"
	"todo/internal/api/todo"

	"github.com/gin-gonic/gin"
)

const (
	groupPath string = "/todo"
	pingPath  string = "/ping"
)

type handler struct{}

func (h *handler) Register(rg *gin.RouterGroup) {
	useCase := todo.GetUseCase()
	group := rg.Group(groupPath)
	{
		group.GET(pingPath, useCase.PongFunction)
	}
}

func GetHandler() api.IHandler {
	return &handler{}
}
