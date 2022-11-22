package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (c *HomeController) HomePageGET(ctx *gin.Context) {
	vals := ctx.GetStringMap("values")
	ctx.HTML(http.StatusOK, "home/index", vals)
}
