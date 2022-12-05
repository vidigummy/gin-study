package handler

import (
	"github.com/dev-yakuza/study-golang/gin/start/github"
	"github.com/gin-gonic/gin"
)

func GetUserUsingLanguageFromGithub(c *gin.Context) {
	userName := c.Param("username")
	userLanguageMap, err := github.GetUserInfo(userName)
	if err != nil {
		panic(err)
	}
	c.JSON(200, userLanguageMap)
}
