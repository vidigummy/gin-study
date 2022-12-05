package main

import (
	"github.com/dev-yakuza/study-golang/gin/start/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
