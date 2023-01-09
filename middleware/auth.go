package middleware

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if session.Get("user") == nil {
			log.Println("User not logged in!")
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
		}
		ctx.Next()
	}
}

func Guest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if session.Get("user") != nil {
			log.Println("User must not be logged in")
			ctx.Redirect(http.StatusFound, "/")
			ctx.Abort()
		}
		ctx.Next()
	}
}

func IsLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		vals := ctx.GetStringMap("values")
		vals["isLoggedIn"] = session.Get("user") != nil
		ctx.Set("values", vals)
		ctx.Next()
	}
}
