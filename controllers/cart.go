package controllers

import (
	"net/http"
	"webmarketplace/database"
	"webmarketplace/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CartController struct{}

func (c *CartController) CartPageGET(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")

	db := database.DB()
	session := sessions.Default(ctx)
	email := session.Get("user")

	var user models.User
	if db.Where("email = ?", email).First(&user).Error != nil {
		ctx.Abort()
		return
	}

	var items []models.Item
	if db.Where("purchaser_id = ?", user.ID).Find(&items).Error != nil {
		ctx.Abort()
		return
	}

	vals["items"] = items

	ctx.HTML(http.StatusOK, "cart/index", vals)
}

func (c *CartController) AddToCartPOST(ctx *gin.Context) {
	db := database.DB()

	var item models.Item
	if db.First(&item, ctx.Param("id")).Error != nil {
		ctx.Abort()
		return
	}

	session := sessions.Default(ctx)
	email := session.Get("user")

	var user models.User
	if db.Where("email = ?", email).First(&user).Error != nil {
		ctx.Abort()
		return
	}

	item.Purchaser = user

	if db.Save(item).Error != nil {
		ctx.Abort()
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/items/")
}

func (c *CartController) RemoveFromCartPOST(ctx *gin.Context) {
	db := database.DB()

	var item models.Item
	if db.First(&item, ctx.Param("id")).Error != nil {
		ctx.Abort()
		return
	}

	session := sessions.Default(ctx)
	email := session.Get("user")

	var user models.User
	if db.Where("email = ?", email).First(&user).Error != nil {
		ctx.Abort()
		return
	}

	if *item.PurchaserID != user.ID {
		ctx.Abort()
		return
	}

	item.PurchaserID = nil

	if db.Save(item).Error != nil {
		ctx.Abort()
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/cart")
}
