package routes

import (
	"webmarketplace/controllers"

	"github.com/gin-gonic/gin"
)

func SetHomeRoutes(r *gin.RouterGroup) {
	c := new(controllers.HomeController)
	r.GET("/", c.HomePageGET)
}
