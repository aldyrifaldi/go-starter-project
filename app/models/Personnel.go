package model

import (
	"time"

	"gorm.io/gorm"
)

type Personnel struct {
	gorm.Model
	UserID uint 
	name string `gorm:"type:varchar(255) not null"`
	phoneNumber int64 `gorm:"type:bigint not null"`
	gender int `gorm:"type:tinyint"`
	religion int `gorm:"type:tinyint"`
	bloodType string `gorm:"type:varchar(255)"`
	birtDate time.Time `gorm:"type:date not null"`
}