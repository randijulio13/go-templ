package app

import (
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/views"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return views.Render(c, views.Home())
	})
}
