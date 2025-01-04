package models

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	UserID   int    `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
