package handler

import "github.com/gin-gonic/gin"

func SayHi(c *gin.Context) {
	c.String(200, "Hello World!")
}
