package middleware

import (
	"fmt"
	"net/http"

	"github.com/dev-yakuza/study-golang/gin/start/models"
	"github.com/gin-gonic/gin"
)

func CheckAuthenticationWithHeader(c *gin.Context) {
	Token := c.GetHeader("access_token")
	fmt.Println(Token)

	if Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "No Auth Token"})
		c.Abort()
		return
	}
	err := models.VerifyToken(Token)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	}
	c.Next()
}
