package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/utils"
	"github.com/randijulio13/go-templ/views"
)

type HomeHandler struct{}

func (h *HomeHandler) Index(c echo.Context) error {
	utils.SetFlashmessages(c, "success", "Flash Message Example")
	utils.SetSession(c, "message", "This Message is From Index")
	return views.Render(c, views.Home(utils.GetFlashmessages(c, "success"), utils.GetFlashmessages(c, "error")))
}

func (h *HomeHandler) About(c echo.Context) error {
	var msg []string
	if msgSession := utils.GetSession(c, "message"); msgSession != nil {
		msg = []string{msgSession.(string), msgSession.(string)}
	}

	utils.DestroySession(c, "message")
	err := utils.GetFlashmessages(c, "error")
	// err := []string{"Test Message1", "Test Message2"}
	return views.Render(c, views.About(msg, err))
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}
