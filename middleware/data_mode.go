package middleware

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DataMode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		mode := session.Get("save")
		log.Println(mode)
		if mode == nil {
			session.Set("save", "off")
			mode = "off"
		}
		vals := ctx.GetStringMap("values")
		vals["saveData"] = mode == "on"
		log.Println(vals["saveData"])
		ctx.Set("values", vals)
		ctx.Next()
	}
}
