package router

import (
	"github.com/dev-yakuza/study-golang/gin/start/github"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	root := router.Group("/")
	{
		root.GET("/", github.GetUserInfo)
	}
	return router
}
