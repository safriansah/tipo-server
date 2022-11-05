package app

import (
	"net/http"
)

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods(http.MethodGet)

	post := a.Router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/api/check-word", a.CheckWordHandler())
	post.Use(checkToken)

	a.Router.HandleFunc("/api/google/login", a.GoToGoogleLoginPage()).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/google/callback", a.GoogleLoginCallback()).Methods(http.MethodGet)
}
