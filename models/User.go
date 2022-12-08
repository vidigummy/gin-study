package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dev-yakuza/study-golang/gin/start/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id            int64 `json:"id" gorm:"primary_key"`
	UserName      string
	UserPassword  string
	UserLanguages []UserLanguage `gorm:"foreignKey:UserIdx"`
}

func SetUser(username string, pw string) *User {
	newUser := User{UserName: username, UserPassword: pw}
	return &newUser
}

func (u *User) UserToString() map[string]string {
	var result = make(map[string]string)
	result["Id"] = strconv.FormatInt(u.Id, 10)
	result["UserName"] = u.UserName
	result["UserPassword"] = u.UserPassword
	return result
}

func (u *User) CommitUser() error {
	db := database.Database()
	result := db.Create(&User{UserName: u.UserName, UserPassword: u.UserPassword})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserFromName(userName string) (*User, error) {
	db := database.Database()
	var result User
	db.First(&result, "user_name = ?", userName)
	fmt.Println(userName, result.CreatedAt)
	if result.UserName != "" {
		return &result, nil
	}
	return nil, errors.New("No Result")
}
