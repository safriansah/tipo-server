package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"type:VARCHAR;NOT NULL" json:"name"`
	Username  string     `gorm:"type:VARCHAR;NOT NULL,unique" json:"username"`
	Email     string     `gorm:"type:VARCHAR;NOT NULL,unique" json:"email"`
	Phone     string     `gorm:"type:VARCHAR;NOT NULL" json:"phone"`
	Picture   string     `gorm:"type:VARCHAR;NOT NULL" json:"picture"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
