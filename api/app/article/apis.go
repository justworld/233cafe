package article

import (
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
		appG.ErrResponse(app.ERROR, err.Error())
	} else {
		s := c.Query("s")
		data, err := GetArticles(p, config.AppSetting.PageSize, s)
		if err != nil {
			appG.ErrResponse(app.ERROR, err.Error())
		} else {

			result := make([] map[string]interface{}, len(data))
			for i, a := range data {
				result[i] = map[string]interface{}{
					"id":          a.ID,
					"category_id": a.ID,
					"title": a.Title,
					"desc": a.Desc,
					"cover": a.Cover,
					"read": a.ReadNum,
					"love": a.LikeNum,
				}
			}

			appG.Response(result)
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
	appG.Response(0)
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
	appG.Response(0)
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
	appG.Response(0)
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
	appG.Response(0)
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
	appG.Response(0)
}