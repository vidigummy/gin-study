package models

import (
	"github.com/dev-yakuza/study-golang/gin/start/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	idx       int64  `json:"idx" gorm:"primary_key"`
	user_name string `json:"user_name" gorm:"-:all", "column:user_name"`
	email     string `json:"email" gorm:"-:all", "column:email"`
}

func SetUser(username string, email string) *User {
	newUser := User{user_name: username, email: email}
	return &newUser
}

func CommitUser(u *User) error {
	db := database.Database()
	result := db.Create(&User{user_name: u.user_name, email: u.email})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
