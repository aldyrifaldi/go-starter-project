package routes

import (
	"net/http"
	"starter_project2/app/controllers"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
)


func Routes() *echo.Echo {
	e := echo.New()

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	
	// //CORS
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK,"Welcome to go starter project")
	})
	e.POST("/register", controllers.Register)

	return e
}