package routes

import (
	"webmarketplace/controllers"
	"webmarketplace/middleware"

	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(r *gin.RouterGroup) {
	c := new(controllers.AuthController)

	guest := r.Group("/")
	guest.Use(middleware.Guest())
	{
		guest.GET("/login", c.LoginPageGET)
		guest.GET("/register", c.RegisterPageGET)
		guest.POST("/login", c.LoginPOST)
		guest.POST("/register", c.RegisterPOST)
	}
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
	{
		auth.POST("/logout", c.LogoutPOST)
	}
}
