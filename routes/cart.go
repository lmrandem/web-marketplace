package routes

import (
	"webmarketplace/controllers"
	"webmarketplace/middleware"

	"github.com/gin-gonic/gin"
)

func SetCartRoutes(r *gin.RouterGroup) {
	c := new(controllers.CartController)
	g := r.Group("/cart")
	g.Use(middleware.AuthRequired())
	{
		g.GET("/", c.CartPageGET)
		g.POST("/add/:id", c.AddToCartPOST)
		g.POST("/remove/:id", c.RemoveFromCartPOST)
	}
}
