package models

import (
	"errors"
	"strconv"

	"github.com/dev-yakuza/study-golang/gin/start/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id            int64 `json:"id" gorm:"primary_key"`
	UserName      string
	Email         string
	UserLanguages []UserLanguage `gorm:"foreignKey:UserIdx"`
}

func SetUser(username string, email string) *User {
	newUser := User{UserName: username, Email: email}
	return &newUser
}

func UserToString(u *User) map[string]string {
	var result = make(map[string]string)
	result["Id"] = strconv.FormatInt(u.Id, 10)
	result["UserName"] = u.UserName
	result["Email"] = u.Email
	return result
}

func CommitUser(u *User) error {
	db := database.Database()
	result := db.Create(&User{UserName: u.UserName, Email: u.Email})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserFromName(userName string) (*User, error) {
	db := database.Database()
	var result User
	db.First(&result, "user_name = ?", userName)
	if result.UserName != "" {
		return &result, nil
	}
	return nil, errors.New("No Result")
}
