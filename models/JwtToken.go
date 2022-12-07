package models

import "github.com/dgrijalva/jwt-go"

type AuthTokenClaims struct {
	UserID             string   `json:"id"`   // 유저 ID
	Name               string   `json:"name"` // 유저 이름
	Role               []string `json:"role"` // 유저 역할
	jwt.StandardClaims          // 표준 토큰 Claims
}
