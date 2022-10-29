package app

import (
	"tipo-server/app/database"
	"tipo-server/app/models"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.TipoDB
	ENV    *models.Config
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("get")
	a.Router.HandleFunc("/api/check-word", a.CheckWordHandler()).Methods("post")
}
