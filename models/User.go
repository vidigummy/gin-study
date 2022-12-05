package models

import (
	"github.com/dev-yakuza/study-golang/gin/start/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Idx      int64 `json:"idx" gorm:"primary_key"`
	UserName string
	Email    string
}

func SetUser(username string, email string) *User {
	newUser := User{UserName: username, Email: email}
	return &newUser
}

func CommitUser(u *User) error {
	db := database.Database()
	result := db.Create(&User{UserName: u.UserName, Email: u.Email})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
