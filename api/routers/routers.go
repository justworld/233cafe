package routers

import (
	"github.com/gin-gonic/gin"

	"233cafe/app/common"
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
		api.GET("/common/hello", common.GetHello)
		api.GET("/article/hello", article.GetHello)
	}

	return r
}
