package controllers

import (
	"fmt"
	"net/http"
	"webmarketplace/database"
	"webmarketplace/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AccountController struct{}

func (c *AccountController) AccountPageGET(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")

	session := sessions.Default(ctx)
	email := session.Get("user")

	var user models.User
	db := database.DB()
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	vals["user"] = user

	ctx.HTML(http.StatusOK, "account/index", vals)
}

func (c *AccountController) ChangePasswordPOST(ctx *gin.Context) {
	currentPW := ctx.PostForm("password")
	newPW := ctx.PostForm("newPassword")
	newPWCheck := ctx.PostForm("newPasswordCheck")

	if currentPW == "" || newPW == "" || newPWCheck == "" {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusBadRequest, "account/index", vals)
		return
	}
	if newPW != newPWCheck {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusBadRequest, "account/index", vals)
		return
	}

	var user models.User
	session := sessions.Default(ctx)
	email := session.Get("user")

	db := database.DB()
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPW)); err != nil {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusBadRequest, "account/index", vals)
		return
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(newPW), bcrypt.DefaultCost)
	if err != nil {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusBadRequest, "account/index", vals)
		return
	}

	user.Password = string(pwHash)
	if db.Save(user).Error != nil {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusBadRequest, "account/index", vals)
		return
	}

	fmt.Println("Testest etstetretrerfwefwe")

	ctx.Redirect(http.StatusSeeOther, "/account")
}

func (c *AccountController) DeleteAccountPOST(ctx *gin.Context) {
	pw := ctx.PostForm("password")

	session := sessions.Default(ctx)
	email := session.Get("user")

	var user models.User
	db := database.DB()
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusNotFound, "account/index", vals)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw)); err != nil {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusBadRequest, "account/index", vals)
		return
	}

	result = db.Delete(&user)
	if result.Error != nil {
		vals := ctx.GetStringMap("values")
		ctx.HTML(http.StatusInternalServerError, "account/index", vals)
		return
	}

	session.Delete("user")
	if err := session.Save(); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
