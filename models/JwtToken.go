package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	expirationTime = 5 * time.Minute
	JwtKey         = []byte("JwtToken")
)

type AuthTokenClaims struct {
	Name string `json:"name"` // 유저 이름
	jwt.StandardClaims
}

func (user *LoginUser) GetJwtToken() (string, error) {
	expirationDate := time.Now().Add(expirationTime)
	claims := &AuthTokenClaims{
		Name: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
		},
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		println(err.Error())
		return "", errors.New("Token signing failed")
	}
	return tokenString, nil
}
