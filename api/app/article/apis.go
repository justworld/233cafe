package article

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"233cafe/app"
	"233cafe/core/api"
)


func GetHello(c *gin.Context) {
	appG := api.Gin{C: c}
	count, err := GetArticles()
	if err != nil {
		msg := "fail"
		appG.Response(http.StatusOK, app.SUCCESS, msg)
	} else {
		appG.Response(http.StatusOK, app.SUCCESS, count)
	}
}