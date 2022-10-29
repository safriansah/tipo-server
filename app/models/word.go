package models

import "time"

type Word struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Input     string     `gorm:"type:VARCHAR;NOT NULL" json:"input"`
	Result    string     `gorm:"type:VARCHAR;NOT NULL" json:"result"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

type PostWord struct {
	Input string `json:"input"`
}

type TipoReturn struct {
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"message"`
	Input     string `json:"input"`
	Result    string `json:"result"`
}
