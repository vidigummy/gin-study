package main

import (
	"github.com/dev-yakuza/study-golang/gin/start/router"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "ping")
	})
	return r
}

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
