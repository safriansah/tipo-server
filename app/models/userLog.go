package models

import "time"

type UserLog struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserId    uint       `gorm:"type:int;NOT NULL" json:"userId"`
	WordId    uint       `gorm:"type:int;NOT NULL" json:"wordId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
