package middleware

import (
	"github.com/gin-gonic/gin"
)

func ValueMap() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vals := make(map[string]any)

		ctx.Set("values", vals)
		ctx.Next()
	}
}
