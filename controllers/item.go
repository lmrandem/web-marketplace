package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"webmarketplace/database"
	"webmarketplace/models"
	"webmarketplace/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemController struct{}

func (c *ItemController) ListItemsPageGET(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")

	var items []models.Item
	var count int64
	searchTerm := ctx.DefaultQuery("search", "")
	currentPage, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize := 12
	offset := (currentPage - 1) * pageSize
	db := database.DB()
	db.Where("title LIKE ?", "%"+searchTerm+"%").
		Offset(offset).
		Limit(pageSize).
		Find(&items).
		Count(&count)

	vals["search"] = searchTerm
	vals["items"] = items
	vals["pages"] = utils.Iterate(1, int64(math.Ceil(float64(count)/float64(pageSize)))+1)
	vals["current"] = currentPage

	log.Println("Test render items")
	ctx.HTML(http.StatusOK, "items/index", vals)
}

func (c *ItemController) ItemPageGET(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")

	id := ctx.Param("id")
	var item models.Item
	db := database.DB()
	db.First(&item, id)

	vals["title"] = item.Title
	vals["item"] = item

	ctx.HTML(http.StatusOK, "items/show", vals)
}

func (c *ItemController) CreateItemPageGET(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")
	ctx.HTML(http.StatusOK, "items/create", vals)
}

func (c *ItemController) CreateItemPOST(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")

	title := ctx.PostForm("title")
	desc := ctx.PostForm("description")
	price := ctx.PostForm("price")
	img, err := ctx.FormFile("image")

	if err != nil {
		ctx.HTML(http.StatusBadRequest, "items/create", vals)
		return
	}

	parts := strings.Split(img.Filename, ".")
	ext := parts[len(parts)-1]

	img.Filename = uuid.New().String() + "." + ext

	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "items/create", vals)
		return
	}

	item := models.Item{Title: title, Description: desc, Price: p, ImageUrl: img.Filename}
	db := database.DB()

	if err := db.Create(&item).Error; err != nil {
		ctx.HTML(http.StatusInternalServerError, "items/create", vals)
		return
	}

	filename := "uploads/" + filepath.Base(img.Filename)
	if err := ctx.SaveUploadedFile(img, filename); err != nil {
		fmt.Println("Error", err)
		ctx.HTML(http.StatusInternalServerError, "items/create", vals)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/items/"+fmt.Sprint(item.ID))
}
