package main

import (
	"os"
	"webmarketplace/database"
	"webmarketplace/middleware"
	"webmarketplace/models"
	"webmarketplace/routes"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	gormsession "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupDatabase() {
	db = database.DB()
	db.AutoMigrate(&models.User{}, &models.Item{})

}

func setupRoutes(r *gin.Engine) {

	g := r.Group("/")

	routes.SetAccountRoutes(g)
	routes.SetAuthRoutes(g)
	routes.SetCartRoutes(g)
	routes.SetHomeRoutes(g)
	routes.SetItemRoutes(g)
}

func main() {
	os.Mkdir("./public/uploads", 0777)
	setupDatabase()
	r := gin.Default()
	r.MaxMultipartMemory = 2 << 20
	store := gormsession.NewStore(db, true, []byte("secret"))
	r.Use(sessions.Sessions("sessid", store))

	// Middleware
	r.Use(middleware.ValueMap())
	// r.Use(middleware.HostValue())
	r.Use(middleware.IsLoggedIn())

	setupRoutes(r)

	r.Static("/assets", "./public/dist")
	r.Static("/uploads", "./uploads")
	r.HTMLRender = ginview.Default()

	r.Run(":8080")
}
