package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"tipo-server/app/utils"
)

func checkToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header["Authorization"]
		log.Printf("authorization::%v", authorization)
		if authorization == nil {
			sendResponse(w, r, nil, http.StatusUnauthorized)
			return
		}
		token := strings.Split(authorization[0], " ")[1]
		log.Printf("token::%v", token)
		status, user := utils.CheckJWTToken(&token)
		if status != http.StatusOK {
			sendResponse(w, r, nil, status)
			return
		}
		log.Printf("user::%v", user)
		r.Header.Set("user-id", fmt.Sprint(user.Id))
		r.Header.Set("username", user.Username)
		r.Header.Set("email", user.Email)
		h.ServeHTTP(w, r)
	})
}
