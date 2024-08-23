package router

import (
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/handler"
)

func SetupRoutes(e *echo.Echo, hh handler.HomeHandlerInterface) {
	e.Static("/assets", "assets")
	e.GET("/", hh.Index)
	e.GET("/about", hh.About)
}
