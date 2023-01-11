package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
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
	routes.SetDataModeRoutes(g)
}

func serveFile(name string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(fmt.Sprintf("./public/dist/%s", name))
		ctx.File(fmt.Sprintf("./public/dist/%s", name))
	}
}

func setupAssetsRoute(r *gin.Engine) {
	mime.AddExtensionType(".js", "application/javascript")
	assets, _ := ioutil.ReadDir("./public/dist")
	for _, asset := range assets {
		log.Println("ASSET NAME:", asset.Name())
		if !asset.IsDir() {
			r.GET(fmt.Sprintf("/%s", asset.Name()), serveFile(asset.Name()))
		}
	}
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
	r.Use(middleware.Assets())
	r.Use(middleware.DataMode())
	r.Use(middleware.IsLoggedIn())

	// Routes
	setupRoutes(r)
	setupAssetsRoute(r)
	r.Static("/uploads", "./uploads")

	r.HTMLRender = ginview.Default()

	r.Run(":8080")
}
