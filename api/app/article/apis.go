package article

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"233cafe/app"
	"233cafe/core/api"
)

// @Summary 获取文章列表
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetAricles(c *gin.Context) {
	appG := api.Gin{C: c}
	p := 1
	limit := 1
	s := ""
	count, err := GetAricles(p, limit, s)
	if err != nil {
		appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	} else {
		appG.Response(http.StatusOK, app.SUCCESS, count)
	}
}

// @Summary 获取文章详情
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetInfo(c *gin.Context) {
	appG := api.Gin{C: c}
	id := 1
	count, err := GetAricle(id)
	if err != nil {
		appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	} else {
		appG.Response(http.StatusOK, app.SUCCESS, count)
	}
}

// @Summary 获取热门阅读
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetHots(c *gin.Context)  {
	appG := api.Gin{C: c}
	count, err := GetHots()
	if err != nil {
		appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	} else {
		appG.Response(http.StatusOK, app.SUCCESS, count)
	}
}

// @Summary 获取推荐阅读
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetRecommends(c *gin.Context)  {
	appG := api.Gin{C: c}
	count, err := GetRecommends()
	if err != nil {
		appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	} else {
		appG.Response(http.StatusOK, app.SUCCESS, count)
	}
}

// @Summary 获取广告
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetAds(c *gin.Context)  {
	appG := api.Gin{C: c}
	count, err := GetAds()
	if err != nil {
		appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	} else {
		appG.Response(http.StatusOK, app.SUCCESS, count)
	}
}

// @Summary 点赞
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func UpVote(c *gin.Context)  {
	appG := api.Gin{C: c}
	count, err := UpVote()
	if err != nil {
		appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	} else {
		appG.Response(http.StatusOK, app.SUCCESS, count)
	}
}