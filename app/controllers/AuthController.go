package controllers

import (
	"fmt"
	"net/http"
	model "starter_project2/app/models"

	// model "starter_project2/app/models"

	"github.com/labstack/echo"
)

func Register(c echo.Context) error{
	fmt.Println(c.FormValue("username"))
	user := model.CreateUser(c.Request().MultipartForm)
	// model := model.CreateUser()
	return c.JSON(http.StatusCreated,user)
}