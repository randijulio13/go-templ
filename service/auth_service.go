package service

import (
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/model"
	"github.com/randijulio13/go-templ/utils"
)

type AuthService struct{}

func (as *AuthService) CreateUser(c echo.Context, u *model.User) error {
	utils.SetFlashmessages(c, "success", "Session message")
	return nil
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
