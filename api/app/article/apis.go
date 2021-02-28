package article

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"233cafe/app"
	"233cafe/config"
	"233cafe/core/api"
)

// @Summary 获取文章列表
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetAricleList(c *gin.Context) {
	appG := api.Gin{C: c}
	p, err := strconv.Atoi(c.Query("p"))
	if err != nil {
		appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	} else {
		s := c.Query("s")
		count, err := GetArticles(p, config.AppSetting.PageSize, s)
		if err != nil {
			appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
		} else {
			appG.Response(http.StatusOK, app.SUCCESS, count)
		}
	}
}

// @Summary 获取文章详情
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetInfo(c *gin.Context) {
	appG := api.Gin{C: c}
	//id := 1
	//count, err := GetArticles()
	//if err != nil {
	//	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	//} else {
	//	appG.Response(http.StatusOK, app.SUCCESS, count)
	//}
	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
}

// @Summary 获取热门阅读
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetHots(c *gin.Context)  {
	appG := api.Gin{C: c}
	//count, err := GetArticles()
	//if err != nil {
	//	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	//} else {
	//	appG.Response(http.StatusOK, app.SUCCESS, count)
	//}
	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
}

// @Summary 获取推荐阅读
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetRecommends(c *gin.Context)  {
	appG := api.Gin{C: c}
	//count, err := GetArticles()
	//if err != nil {
	//	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	//} else {
	//	appG.Response(http.StatusOK, app.SUCCESS, count)
	//}
	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
}

// @Summary 获取广告
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetAds(c *gin.Context)  {
	appG := api.Gin{C: c}
	//count, err := GetArticles()
	//if err != nil {
	//	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	//} else {
	//	appG.Response(http.StatusOK, app.SUCCESS, count)
	//}
	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
}

// @Summary 点赞
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func UpVote(c *gin.Context)  {
	appG := api.Gin{C: c}
	//count, err := GetArticles()
	//if err != nil {
	//	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
	//} else {
	//	appG.Response(http.StatusOK, app.SUCCESS, count)
	//}
	appG.Response(http.StatusOK, app.SUCCESS, app.ERROR)
}