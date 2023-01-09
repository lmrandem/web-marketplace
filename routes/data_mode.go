package routes

import (
	"webmarketplace/controllers"

	"github.com/gin-gonic/gin"
)

func SetDataModeRoutes(r *gin.RouterGroup) {
	c := new(controllers.DataModeController)
	r.POST("/mode", c.SetDataModePOST)
}
