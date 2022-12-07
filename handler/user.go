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
	newUserPassword := c.Param("password")
	newUser := models.SetUser(newUserName, newUserPassword)
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

func LoginUser(c *gin.Context) {
	req := &models.LoginUser{}
	err := c.Bind(req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	fmt.Println(req.UserName)
	nowUserName := req.UserName
	user, err := models.GetUserFromName(nowUserName)
	if err != nil {
		c.JSON(403, err.Error())
		c.Abort()
	}
	if user.UserPassword != req.Password {
		c.JSON(403, err.Error())
		c.Abort()
	}

	c.Status(200)

}
