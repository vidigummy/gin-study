package router

import (
	"github.com/dev-yakuza/study-golang/gin/start/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	root := router.Group("/")
	{
		root.GET("", handler.SayHi)
	}
	user := router.Group("/user")
	{
		user.GET("/:username", handler.GetUserUsingLanguageFromGithub)
		user.GET("/new/:username/:email", handler.ResistNewUser)
		user.GET("/find/:username", handler.FindUserWithName)
	}
	return router
}
