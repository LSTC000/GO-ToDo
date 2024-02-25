package intreface

import "github.com/gin-gonic/gin"

type IToDoUseCase interface {
	PongFunction(ctx *gin.Context)
}
