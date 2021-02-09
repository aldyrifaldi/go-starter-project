package databases

import (
	"fmt"
	model "starter_project2/app/models"
	"starter_project2/config"
)


func AutoMigrate() {
	user := model.User{}
	personnel := model.Personnel{}
	db := config.Database()
	if db != nil {
		fmt.Println(db)
	}
	db.AutoMigrate(user,personnel)
}

func MigrateUp() {
	user := model.User{}
	personnel := model.Personnel{}
	db := config.Database()
	if db != nil {
		fmt.Println(db)
	}
	db.Migrator().CreateTable(user,personnel)
}

func MigrateDown() {
	user := model.User{}
	personnel := model.Personnel{}
	db := config.Database()
	if db != nil {
		fmt.Println(db)
	}
	
	db.Migrator().DropTable(user,personnel)
}