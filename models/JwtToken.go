package models

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var (
	expirationTime = 5 * time.Minute
)

type AuthTokenClaims struct {
	Name string `json:"name"` // 유저 이름
	jwt.StandardClaims
}

func (user *LoginUser) GetJwtToken() (string, error) {
	JwtKeyString, _ := getJwtSecretFromEnv()
	JwtKey := []byte(JwtKeyString)
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

func VerifyToken(token string) error {
	claims := &AuthTokenClaims{}
	JwtKeyString, _ := getJwtSecretFromEnv()
	JwtKey := []byte(JwtKeyString)
	result, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)
	return nil
}

func getJwtSecretFromEnv() (string, error) {
	err := godotenv.Load("dev.env")
	if err != nil {
		fmt.Println(err.Error())
	}
	return os.Getenv("JWT_TOKEN_SECRET"), nil
}
