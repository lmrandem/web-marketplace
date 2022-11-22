package routes

import (
	"webmarketplace/controllers"
	"webmarketplace/middleware"

	"github.com/gin-gonic/gin"
)

func SetItemRoutes(r *gin.RouterGroup) {
	c := new(controllers.ItemController)
	g := r.Group("/items")
	{
		g.GET("/", c.ListItemsPageGET)
		g.GET("/:id", c.ItemPageGET)
	}
	auth := r.Group("/items")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/create", c.CreateItemPageGET)
		auth.POST("/create", c.CreateItemPOST)
	}
}
