package utils

import "github.com/gin-gonic/gin"

func RespondWithJSON(ctx *gin.Context, code int, payload interface{}) {
	ctx.JSON(code, payload)
}
