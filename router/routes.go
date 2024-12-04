package router

import (
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/handler"
)

func SetupRoutes(e *echo.Echo, hh *handler.HomeHandler, ah *handler.AuthHandler) {
	e.Static("/assets", "assets")
	e.GET("/", hh.Index)
	e.GET("/about", hh.About)

	e.GET("/register", ah.CreateUser)
	e.POST("/register", ah.CreateUser)
}
