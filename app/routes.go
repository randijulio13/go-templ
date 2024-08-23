package app

import (
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/handler"
	"github.com/randijulio13/go-templ/router"
)

func Routes(e *echo.Echo) {
	hh := handler.NewHomeHandler()

	router.SetupRoutes(e, hh)
}
