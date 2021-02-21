package common

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"233cafe/app"
	"233cafe/core/api"
)


func GetHello(c *gin.Context) {
	appG := api.Gin{C: c}
	msg := "hello"
	appG.Response(http.StatusOK, app.SUCCESS, msg)
}