package v1

import (
	"todo/internal/intreface"
	"todo/internal/todo"

	"github.com/gin-gonic/gin"
)

const (
	groupPath string = "/todo"
	pingPath  string = "/ping"
)

type handler struct{}

func GetHandler() intreface.IHandler {
	return &handler{}
}

func (h *handler) Register(rg *gin.RouterGroup) {
	todoUseCase := todo.GetUseCase()
	todoGroup := rg.Group(groupPath)
	{
		todoGroup.GET(pingPath, todoUseCase.PongFunction)
	}
}
