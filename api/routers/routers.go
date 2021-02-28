package routers

import (
	"github.com/gin-gonic/gin"

	"233cafe/app/article"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.Use()
	{
		api.GET("/article", article.GetAricleList)
		api.GET("/hot", article.GetHots)
		api.GET("/article/:id", article.GetInfo)
		api.GET("/article/:id/recommend", article.GetRecommends)
		api.GET("/article/:id/ad", article.GetAds)
		api.GET("/article/:id/love", article.UpVote)
	}

	return r
}
