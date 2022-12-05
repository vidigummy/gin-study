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

// func database() {
// 	datetimePrecision := 2
// 	db, err := gorm.Open(mysql.New(mysql.Config{
// 		DSN:                       "vidi:vidi@tcp(localhost:3306)/RPG?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
// 		DefaultStringSize:         256,                                                                       // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
// 		DisableDatetimePrecision:  true,                                                                      // disable datetime precision support, which not supported before MySQL 5.6
// 		DefaultDatetimePrecision:  &datetimePrecision,                                                        // default datetime precision
// 		DontSupportRenameIndex:    true,                                                                      // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
// 		DontSupportRenameColumn:   true,                                                                      // use change when rename column, rename rename not supported before MySQL 8, MariaDB
// 		SkipInitializeWithVersion: false,                                                                     // smart configure based on used version
// 	}), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(db.Config)
// }
