package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuthenticationWithHeader(c *gin.Context) {
	Token := c.GetHeader("access_token")
	fmt.Println(Token)

	if Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "No AUth Token"})
		c.Abort()
		return
	}

	c.Next()
}
