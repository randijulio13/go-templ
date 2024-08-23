package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/views"
)

type HomeHandlerInterface interface {
	Index(echo.Context) error
	About(echo.Context) error
}

type HomeHandler struct {
}

func (h *HomeHandler) Index(c echo.Context) error {
	return views.Render(c, views.Home())
}

func (h *HomeHandler) About(c echo.Context) error {
	return views.Render(c, views.About())
}

func NewHomeHandler() HomeHandlerInterface {
	return &HomeHandler{}
}
