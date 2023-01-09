package middleware

import "github.com/gin-gonic/gin"

func CacheControl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "GET" || ctx.Request.Method == "" {
			ctx.Header("Cache-Control", "max-age=31536000, immutable")
		}
		ctx.Next()
	}
}
