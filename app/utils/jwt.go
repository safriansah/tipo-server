package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"tipo-server/app/models"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var jwtKey = []byte("")

func SetJWTKey() {
	jwtKey = []byte(viper.GetString("JWT_KEY"))
}

func CreateJWTToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &models.JWTClaims{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func CheckJWTToken(token *string) (int, *models.JWTClaims) {
	claims := &models.JWTClaims{}
	result, err := jwt.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if fmt.Sprint(err) == fmt.Sprint(jwt.ErrSignatureInvalid) {
			log.Print("masuk err")
			return http.StatusUnauthorized, nil
		}
		return http.StatusBadRequest, nil
	}
	if !result.Valid {
		return http.StatusUnauthorized, nil
	}
	return http.StatusOK, claims
}
