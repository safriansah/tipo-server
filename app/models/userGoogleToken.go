package models

import (
	"time"
)

type UserGoogleToken struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	UserId       uint       `gorm:"type:int;NOT NULL" json:"userId"`
	GoogleId     string     `gorm:"type:varchar;NOT NULL" json:"googleId"`
	AccessToken  string     `gorm:"type:varchar;NOT NULL" json:"access_token"`
	TokenType    string     `gorm:"type:varchar;NOT NULL" json:"token_type,omitempty"`
	RefreshToken string     `gorm:"type:varchar;NOT NULL" json:"refresh_token,omitempty"`
	Expiry       time.Time  `json:"expiry,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
}
