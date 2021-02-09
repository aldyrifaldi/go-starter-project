package main

import (
	"fmt"
	"os"
	"starter_project2/databases"
	"starter_project2/routes"
	"strings"
)

func main() {
	
	console()

	e := routes.Routes()
	e.Start(":9200")
}


func console() {
	argument := os.Args[1:]
	order := strings.Join(argument," ")
	switch order {
		case "gorm migrate up": 
			databases.MigrateUp()
			
		case "gorm migrate down":	
			databases.MigrateDown()
		case "gorm migrate automigrate":
			databases.AutoMigrate()
		default:	
			fmt.Println("Welcome to gorm executable command")
			fmt.Println("Usage:")
			fmt.Println("gorm <command> [argument]")
			fmt.Println("")
			fmt.Println("available command :")
			fmt.Println("     - migrate")
			fmt.Println("")
			fmt.Println("available argument:")
			fmt.Println("     - up")
			fmt.Println("     - down")
			fmt.Println("     - automigrate")
	}
}