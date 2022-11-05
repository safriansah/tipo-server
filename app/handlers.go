package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"tipo-server/app/clients"
	"tipo-server/app/models"

	"github.com/jinzhu/gorm"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}
}

func (a *App) CheckWordHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostWord{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("cannot parse body, err=%v\n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		word, err := a.DB.FindWordByInput(&req.Input)
		if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
			log.Printf("error find word, err=%v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		if word.ID != 0 {
			sendResponse(w, r, word, http.StatusOK)
			return
		}

		result, err := clients.FetchCheckTypo(req.Input)
		if err != nil {
			log.Printf("cannot parse body, err=%v\n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		newWord := &models.Word{
			ID:     0,
			Input:  req.Input,
			Result: *result,
		}

		newWord, err = a.DB.CreateWord(newWord)
		if err != nil {
			log.Printf("cannot save word in db, err=%v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		sendResponse(w, r, newWord, http.StatusOK)
	}
}

func (a *App) GoToGoogleLoginPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := clients.GetGoogleLoginUrl()
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func (a *App) GoogleLoginCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		code := r.FormValue("code")

		if state != clients.GetGoogleOauthState() {
			http.Redirect(w, r, "/api/google/invalid", http.StatusTemporaryRedirect)
			return
		}

		data, err := clients.HandleGoogleCodeExcange(code)
		if err != nil {
			log.Printf("HandleGoogleCodeExcange, err=%v\n", err)
			http.Redirect(w, r, "/api/google/invalid", http.StatusTemporaryRedirect)
			return
		}

		user, err := a.DB.FindUserByEmail(&data.Email)
		if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
			log.Printf("a.DB.FindUserByEmail, err=%v\n", err)
			http.Redirect(w, r, "/api/google/error", http.StatusTemporaryRedirect)
			return
		}
		if user.ID == 0 {
			user = &models.User{
				ID:       0,
				Name:     data.Name,
				Username: strings.Split(data.Email, "@")[0],
				Email:    data.Email,
				Picture:  data.Picture,
			}

			user, err = a.DB.SaveUser(user)
			if err != nil {
				log.Printf("a.DB.SaveUser, err=%v\n", err)
				http.Redirect(w, r, "/api/google/error", http.StatusTemporaryRedirect)
				return
			}
		}

		googleToken, err := a.DB.FindGoogleTokenByUserId(getUintPointer(user.ID))
		if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
			log.Printf("a.DB.FindGoogleTokenByUserId, err=%v\n", err)
			http.Redirect(w, r, "/api/google/error", http.StatusTemporaryRedirect)
			return
		}
		if googleToken.ID != 0 {
			googleToken.AccessToken = data.AccessToken
			googleToken.RefreshToken = data.RefreshToken
			googleToken.TokenType = data.TokenType
			googleToken.Expiry = data.Expiry
			err = a.DB.UpdateGoogleToken(googleToken)
			if err != nil {
				http.Redirect(w, r, "/api/google/error", http.StatusTemporaryRedirect)
				return
			}
		} else {
			googleToken = &models.UserGoogleToken{
				ID:           0,
				UserId:       user.ID,
				GoogleId:     data.Id,
				AccessToken:  data.AccessToken,
				RefreshToken: data.RefreshToken,
				TokenType:    data.TokenType,
				Expiry:       data.Expiry,
			}
			googleToken, err = a.DB.SaveUserGoogleToken(googleToken)
			if err != nil {
				http.Redirect(w, r, "/api/google/error", http.StatusTemporaryRedirect)
				return
			}
		}

		log.Printf("googleToken::%v", googleToken)
		sendResponse(w, r, user, http.StatusOK)
	}
}
