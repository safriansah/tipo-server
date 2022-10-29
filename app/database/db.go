package database

import (
	"fmt"
	"log"
	"tipo-server/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type TipoDB interface {
	Open(config models.Config) error
	Close() error
	CreateWord(*models.Word) (*models.Word, error)
	FindWordByInput(*string) (*models.Word, error)
}

type DB struct {
	db *gorm.DB
}

func (d *DB) Open(config models.Config) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_SSL)
	pg, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	log.Println("connect to database")
	d.db = pg
	d.db.LogMode(true)

	var word models.Word
	pg.AutoMigrate(&word)
	log.Println("run migration")

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
