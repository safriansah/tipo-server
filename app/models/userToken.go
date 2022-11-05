package models

import "time"

type UserToken struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserId    uint       `gorm:"type:int;NOT NULL" json:"userId"`
	Token     string     `gorm:"type:varchar;NOT NULL" json:"token"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
