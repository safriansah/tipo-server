package main

import (
	"log"
	"net/http"
	"os"
	"tipo-server/app"
	"tipo-server/app/database"
	"tipo-server/app/models"

	"github.com/spf13/viper"
)

func main() {
	app := app.New()

	config, err := LoadConfig(".")
	check(err)
	app.ENV = &config

	port := config.PORT

	app.DB = &database.DB{}
	err = app.DB.Open(config)
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

func LoadConfig(path string) (config models.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// https://dev.to/lucasnevespereira/write-a-rest-api-in-golang-following-best-practices-pe9
