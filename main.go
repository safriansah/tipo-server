package main

import (
	"log"
	"net/http"
	"os"
	"tipo-server/app"
	"tipo-server/app/clients"
	"tipo-server/app/database"
	"tipo-server/app/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	clients.InitializeOAuthGoogle()

	utils.SetJWTKey()

	app := app.New()
	app.DB = &database.DB{}
	err = app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("app running on " + port)
	err = http.ListenAndServe(":"+port, nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}

// https://dev.to/lucasnevespereira/write-a-rest-api-in-golang-following-best-practices-pe9
