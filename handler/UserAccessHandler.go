package handler

import (
	"errors"
	"fmt"

	"github.com/dev-yakuza/study-golang/gin/start/models"
	"github.com/gin-gonic/gin"
)

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
	}
	newToken, err := req.GetJwtToken()
	if err != nil {
		fmt.Println(err)
	}
	c.SetCookie("access_token", newToken, 1800, "", "", false, false)
	c.Status(200)

}

func SignUp(c *gin.Context) {
	// Body값 확인
	req := &models.LoginUser{}
	err := c.Bind(req)
	if err != nil {
		c.Error(errors.New("Can't Find User Request Body"))
		return
	}
	// 로그인 되는지?
	_, err = models.GetUserFromName(req.UserName)
	if err == nil {
		c.String(409, "중복됩니다~")
	}
	newUser := models.SetUser(req.UserName, req.Password)
	newUser.CommitUser()
	newToken, err := req.GetJwtToken()
	if err != nil || newToken == "" {
		fmt.Println(err)
		c.String(500, "fucked Up")
	}
	fmt.Println("New Token is : ", newToken)
	c.SetCookie("access_token", newToken, 1800, "", "", false, false)
	c.Status(200)
}
