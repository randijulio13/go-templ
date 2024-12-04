package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/model"
	"github.com/randijulio13/go-templ/service"
	"github.com/randijulio13/go-templ/utils"
	"github.com/randijulio13/go-templ/views"
)

type AuthHandler struct {
	authSvc *service.AuthService
}

func (h *AuthHandler) CreateUser(c echo.Context) error {

	if c.Request().Method == "GET" {
		return views.Render(c, views.Register(utils.GetFlashmessages(c, "success"), utils.GetFlashmessages(c, "error")))
	}

	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.authSvc.CreateUser(c, user)

	return c.Redirect(http.StatusSeeOther, "/register")
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authSvc: as,
	}
}
