package utils

import (
	"github.com/gin-gonic/gin"
)

func GetContextValues(ctx *gin.Context) map[string]any {
	return ctx.GetStringMap("values")
}
