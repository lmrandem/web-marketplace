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

type AuthController struct{}

func (c *AuthController) LoginPageGET(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "auth/login", gin.H{})
}

func (c *AuthController) RegisterPageGET(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")
	ctx.HTML(http.StatusOK, "auth/register", vals)
}

func (c *AuthController) LoginPOST(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if email == "" || password == "" {
		ctx.HTML(http.StatusBadRequest, "auth/login", vals)
		return
	}

	var user models.User
	db := database.DB()

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		ctx.HTML(http.StatusBadRequest, "auth/login", vals)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "auth/login", vals)
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", email)
	if err := session.Save(); err != nil {
		ctx.HTML(http.StatusBadRequest, "auth/login", vals)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func (c *AuthController) RegisterPOST(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")

	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	address := ctx.PostForm("address")
	postCode := ctx.PostForm("postCode")
	place := ctx.PostForm("place")
	country := ctx.PostForm("country")

	fmt.Println(place)
	fmt.Println(name, email, password, address, postCode, place, country)

	if name == "" || email == "" || password == "" || address == "" ||
		postCode == "" || place == "" || country == "" {
		ctx.HTML(http.StatusBadRequest, "auth/register", vals)
		return
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		ctx.HTML(http.StatusBadRequest, "auth/register", vals)
		return
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(pwHash),
		Address:  address,
		PostCode: postCode,
		Place:    place,
		Country:  country,
	}

	db := database.DB()
	result := db.Create(&user)

	if result.Error != nil {
		ctx.HTML(http.StatusBadRequest, "auth/register", vals)
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", email)
	if err := session.Save(); err != nil {
		ctx.HTML(http.StatusBadRequest, "auth/register", vals)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func (c *AuthController) LogoutPOST(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("user")
	if err := session.Save(); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/")
}
