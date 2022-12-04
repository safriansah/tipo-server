package database

import (
	"fmt"
	"log"
	"os"
	"tipo-server/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	SaveUserToken(*models.UserToken) (*models.UserToken, error)

	SaveUserLog(*models.UserLog) (*models.UserLog, error)
	FindUserLogByUserId(uint) (*[]models.UserLog, error)
}

type DB struct {
	db *gorm.DB
}

func (d *DB) Open() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))
	pg, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	log.Println("connect to database")
	d.db = pg
	d.db.LogMode(true)

	var word models.Word
	pg.AutoMigrate(&word)

	var user models.User
	pg.AutoMigrate(&user)

	var userGoogleToken models.UserGoogleToken
	pg.AutoMigrate(&userGoogleToken)

	var userToken models.UserToken
	pg.AutoMigrate(&userToken)

	var userLog models.UserLog
	pg.AutoMigrate(&userLog)

	log.Println("run migration")

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
