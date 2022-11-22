package routes

import (
	"webmarketplace/controllers"
	"webmarketplace/middleware"

	"github.com/gin-gonic/gin"
)

func SetAccountRoutes(r *gin.RouterGroup) {
	c := new(controllers.AccountController)
	g := r.Group("/account")
	g.Use(middleware.AuthRequired())
	{
		g.GET("/", c.AccountPageGET)
		g.POST("change-password", c.ChangePasswordPOST)
		g.POST("/delete", c.DeleteAccountPOST)
	}
}
