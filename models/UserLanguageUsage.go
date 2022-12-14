package models

import (
	"errors"
	"fmt"

	"github.com/dev-yakuza/study-golang/gin/start/database"
	"gorm.io/gorm"
)

type UserLanguage struct {
	gorm.Model
	Id          int64 `json:"id" gorm:"primary_key"`
	UserIdx     int64
	UseLanguage string
	Amount      int64
}

func SetUserLanguage(userName string, language string, amount int64) *UserLanguage {
	user, err := GetUserFromName(userName)
	if err != nil {
		panic(err)
	}
	newLanguage := UserLanguage{UserIdx: user.Id, UseLanguage: language, Amount: amount}
	return &newLanguage
}

func (l *UserLanguage) CommitUserLanguage() error {
	result, err := GetUserLanguage(l.UserIdx, l.UseLanguage)
	if err != nil {
		// 중복이 없다면.
		fmt.Println("Not duplicate!")
		db := database.Database()
		result := db.Create(&l)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	// 중복이라면
	if result.Amount != l.Amount {
		l.UpdateUserLanguage(result)
		return nil
	}
	return nil
}
func (l *UserLanguage) UpdateUserLanguage(old *UserLanguage) {
	old.Amount = l.Amount
	db := database.Database()
	db.Save(&old)
}
func GetUserLanguage(userIdx int64, language string) (*UserLanguage, error) {

	db := database.Database()
	var result UserLanguage
	db.First(&result, "user_idx = ? AND use_language = ?", userIdx, language)
	if result.UseLanguage != "" {
		return &result, nil
	}
	return nil, errors.New("No Result[UserLanguage]")
}

func SetUserUseLanguagesMap(m map[string]int, userName string) error {
	for key, value := range m {
		newLanguage := SetUserLanguage(userName, key, int64(value))
		err := newLanguage.CommitUserLanguage()
		if err != nil {
			return err
		}
	}
	return nil
}
