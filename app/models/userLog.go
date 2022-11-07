package models

import "time"

type UserLog struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserId    uint       `gorm:"type:int;NOT NULL" json:"userId"`
	User      User       `json:"user"`
	WordId    uint       `gorm:"type:int;NOT NULL" json:"wordId"`
	Word      Word       `json:"word"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
