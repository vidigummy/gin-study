package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SayHiMiddle(c *gin.Context) {
	fmt.Println("Hello World!")
}
