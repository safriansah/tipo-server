package app

import (
	"fmt"
	"log"
	"net/http"
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

func (a *App) GetGoogleLoginUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := clients.GetGoogleLoginUrl()
		if err != nil {
			log.Printf("error get url, err=%v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, models.GoogleUrlResponse{
			Url: url,
		}, http.StatusOK)
	}
}
