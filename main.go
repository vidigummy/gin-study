package main

import (
	_ "github.com/go-sql-driver/mysql"

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
	// database()
	r.Run(":8080")
}

// func database() {
// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s", "vidi", "vidi", "tcp", "localhost", "3306", "TEST"))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	var user User
// 	err = db.QueryRow("SELECT * FROM USER WHERE idx =1").Scan(&user.idx, &user.user_name, &user.email)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(user)
// }
