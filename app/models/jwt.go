package models

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
