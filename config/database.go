package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB{
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "roots:@tcp(127.0.0.1:3306)/go_starter_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}