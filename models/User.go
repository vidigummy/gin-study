package models

import (
	"errors"
	"strconv"

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

func UserToString(u *User) map[string]string {
	var result = make(map[string]string)
	result["Idx"] = strconv.FormatInt(u.Idx, 10)
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
	// db.Model(User{UserName: userName}).First(&result)\
	db.First(&result, "user_name = ?", userName)
	if result.UserName != "" {
		return &result, nil
	}
	return nil, errors.New("No Result")
}
