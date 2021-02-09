package model

import (
	"time"

	"gorm.io/gorm"
)

type Personnel struct {
	gorm.Model
	UserID int 
	FullName string `gorm:"type:varchar(255)"`
	PhoneNumber int64 `gorm:"type:bigint"`
	Gender int `gorm:"type:tinyint"`
	Religion int `gorm:"type:tinyint"`
	BloodType string `gorm:"type:varchar(255)"`
	BirtDate time.Time `gorm:"type:date"`
}