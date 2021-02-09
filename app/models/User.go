package model

import (
	"database/sql"
	"fmt"
	"mime/multipart"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model // auto generate id, created_at, updated_at and deleted_at
	Username     string `gorm:"type:varchar(255) not null"`
	Email        *string `gorm:"type:varchar(255) not null"`
	MemberNumber sql.NullString `gorm:"not null"`
	ActivedAt    sql.NullTime
	Personnel Personnel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}


func CreateUser(request *multipart.Form) error {
	data, err := fmt.Println(request)
	fmt.Println(data)
	return err
}