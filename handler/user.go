package handler

import (
	"fmt"

	"github.com/dev-yakuza/study-golang/gin/start/github"
	"github.com/dev-yakuza/study-golang/gin/start/models"
	"github.com/gin-gonic/gin"
)

func GetUserUsingLanguageFromGithub(c *gin.Context) {
	userName := c.Param("username")
	userLanguageMap, err := github.GetUserInfo(userName)
	if err != nil {
		panic(err)
	}
	err = models.SetUserUseLanguagesMap(userLanguageMap, userName)
	if err != nil {
		c.Status(500)
	}
	c.JSON(200, userLanguageMap)
}

func ResistNewUser(c *gin.Context) {
	newUserName := c.Param("username")
	newUserEmail := c.Param("email")
	newUser := models.SetUser(newUserName, newUserEmail)
	fmt.Println(newUser)
	err := models.CommitUser(newUser)
	if err != nil {
		c.String(500, err.Error())
	}
	c.String(200, "OK")
}

func FindUserWithName(c *gin.Context) {
	userName := c.Param("username")
	user, err := models.GetUserFromName(userName)
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, models.UserToString(user))
}
