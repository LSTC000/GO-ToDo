package todo

import (
	"net/http"
	"todo/internal/intreface"
	"todo/internal/model"

	"github.com/gin-gonic/gin"
)

type useCase struct{}

func GetUseCase() intreface.IToDoUseCase {
	return &useCase{}
}

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
