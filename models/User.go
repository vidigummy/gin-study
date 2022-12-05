package models

type User struct {
	idx       int64  `json:"idx" gorm:"primary_key"`
	user_name string `json:"user_name"`
	email     string `json:"email"`
}

func SetUser(username string, email string) *User {
	newUser := User{user_name: username, email: email}
	return &newUser
}
