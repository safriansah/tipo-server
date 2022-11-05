package database

import (
	"fmt"
	"log"
	"tipo-server/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type TipoDB interface {
	Open() error
	Close() error

	CreateWord(*models.Word) (*models.Word, error)
	FindWordByInput(*string) (*models.Word, error)

	SaveUserGoogleToken(*models.UserGoogleToken) (*models.UserGoogleToken, error)
	FindGoogleTokenByUserId(*uint) (*models.UserGoogleToken, error)
	UpdateGoogleToken(*models.UserGoogleToken) error

	SaveUser(*models.User) (*models.User, error)
	FindUserByEmail(*string) (*models.User, error)
}

type DB struct {
	db *gorm.DB
}

func (d *DB) Open() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", viper.GetString("DB_HOST"), viper.GetInt("DB_PORT"), viper.GetString("DB_USER"), viper.GetString("DB_PASS"), viper.GetString("DB_NAME"), viper.GetString("DB_SSL"))
	pg, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	log.Println("connect to database")
	d.db = pg
	d.db.LogMode(true)

	var word models.Word
	var user models.User
	var userGoogleToken models.UserGoogleToken
	pg.AutoMigrate(&word)
	pg.AutoMigrate(&user)
	pg.AutoMigrate(&userGoogleToken)
	log.Println("run migration")

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
