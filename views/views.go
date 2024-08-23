package views

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(e echo.Context, v templ.Component) error {
	return v.Render(e.Request().Context(), e.Response().Writer)
}
