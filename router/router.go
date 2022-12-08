package router

import (
	"github.com/dev-yakuza/study-golang/gin/start/handler"
	"github.com/dev-yakuza/study-golang/gin/start/middleware"
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
		user.Use(middleware.SayHiMiddle)
		user.GET("/github/:username", handler.GetUserUsingLanguageFromGithub)
		user.GET("/new/:username/:password", handler.ResistNewUser)
		user.GET("/find/:username", handler.FindUserWithName)
		user.POST("/login", handler.LoginUser)
		user.POST("/signup", handler.SignUp)
	}
	authTest := router.Group("/auth")
	{
		authTest.Use(middleware.CheckAuthenticationWithHeader)
		authTest.GET("/", handler.SayHi)

	}
	return router
}
