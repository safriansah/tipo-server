package app

import (
	"net/http"
)

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods(http.MethodGet)

	auth := a.Router.NewRoute().Subrouter()
	auth.Use(checkToken)

	post := auth.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/api/check-word", a.CheckWordHandler())

	get := auth.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/api/my-words", a.GetMyLog())

	a.Router.HandleFunc("/api/google/login", a.GoToGoogleLoginPage()).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/google/callback", a.GoogleLoginCallback()).Methods(http.MethodGet)
}
