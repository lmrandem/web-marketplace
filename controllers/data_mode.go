package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DataModeController struct{}

func (c *DataModeController) SetDataModePOST(ctx *gin.Context) {
	session := sessions.Default(ctx)
	mode := session.Get("save")
	if mode == "on" {
		session.Set("save", "off")
	} else {
		session.Set("save", "on")
	}
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/")
}
