package todo

import (
	"net/http"
	"todo/internal/model"

	"github.com/gin-gonic/gin"
)

type IUseCase interface {
	PongFunction(ctx *gin.Context)
}

type useCase struct{}

// PongFunction
// @Summary Ping
// @Tags ToDo
// @Description Ping Pong
// @Accept json
// @Produce json
// @Success 200 {object} model.Message
// @Router /todo/ping [get]
func (u *useCase) PongFunction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.Message{Message: "pong"})
}

func GetUseCase() IUseCase {
	return &useCase{}
}
