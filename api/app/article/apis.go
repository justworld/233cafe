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
		return
	} else {
		s := c.Query("s")
		category, err := strconv.Atoi(c.Query("cz"))
		if err != nil {
			category = app.CATEGORY_ALL
		}
		data, err := GetArticles(category, p, config.AppSetting.PageSize, s, false, false)
		if err != nil {
			appG.ErrResponse(app.ERROR, err.Error())
			return
		} else {
			result := make([]map[string]interface{}, len(data))
			for i, a := range data {
				result[i] = map[string]interface{}{
					"id":    a.ID,
					"title": a.Title,
					"desc":  a.Desc,
					"cover": a.Cover,
					"read":  a.ReadNum,
					"love":  a.LikeNum,
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.ErrResponse(app.ERROR, err.Error())
		return
	} else {
		instance, err := GetArticleInfo(id)
		if err != nil {
			appG.ErrResponse(app.ERROR, err.Error())
			return
		}

		// 获取推荐
		recommends := make([]interface{}, 4)
		for i := range recommends {
			recommends[i] = map[string]interface{}{
				"id":    1,
				"title": "北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧",
				"cover": "http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg",
			}
		}
		// 获取广告
		ads := make([]map[string]interface{}, 4)
		for i := range ads {
			ads[i] = map[string]interface{}{
				"url":   "https://www.baidu.com",
				"cover": "http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg",
			}
		}

		result := map[string]interface{}{
			"id":         instance.ID,
			"title":      instance.Title,
			"content":    instance.HtmlContent,
			"recommends": recommends,
			"ads":        ads,
		}

		// 增加浏览记录
		AddArticleRead(id)

		appG.Response(result)
	}
}

// @Summary 获取热门阅读
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetHots(c *gin.Context) {
	appG := api.Gin{C: c}
	hots := make([]interface{}, 7)
	for i := range hots {
		hots[i] = map[string]interface{}{
			"id":    1,
			"title": "北阿坎德邦仍有171人失踪，救援队与时间赛跑，印网民：向中国要一些大型钻机吧",
			"cover": "http://www.santaihu.com/e/data/tmp/titlepic/893f1d9c3c1becfc4155e14787255f00.jpg",
		}
	}
	appG.Response(hots)
}

func GetHotWord(c *gin.Context) {
	appG := api.Gin{C: c}
	appG.Response(map[string]interface{}{
		"hot_word": "hot",
	})
}

// @Summary 获取推荐阅读
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
func GetRecommends(c *gin.Context) {
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
func GetAds(c *gin.Context) {
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
func UpVote(c *gin.Context) {
	appG := api.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.ErrResponse(app.ERROR, err.Error())
		return
	} else {
		AddArticleLike(id)
		appG.Response(0)
	}
}
