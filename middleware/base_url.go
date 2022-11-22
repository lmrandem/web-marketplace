package middleware

import "github.com/gin-gonic/gin"

func HostValue() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vals := ctx.GetStringMap("values")
		vals["host"] = ctx.Request.URL.Host
		ctx.Set("values", vals)
		ctx.Next()
	}
}
