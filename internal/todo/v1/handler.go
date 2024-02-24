package v1

import (
	"net/http"
	"todo/internal/intreface"

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

func (h handler) Register(rg *gin.RouterGroup) {
	todo := rg.Group(groupPath)
	{
		todo.GET(pingPath, pongFunction)
	}
}

func pongFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
