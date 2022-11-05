package app

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("get")

	a.Router.HandleFunc("/api/check-word", a.CheckWordHandler()).Methods("post")

	a.Router.HandleFunc("/api/google/login", a.GoToGoogleLoginPage()).Methods("get")
	a.Router.HandleFunc("/api/google/callback", a.GoogleLoginCallback()).Methods("get")
}
