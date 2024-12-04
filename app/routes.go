package app

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/randijulio13/go-templ/handler"
	"github.com/randijulio13/go-templ/router"
	"github.com/randijulio13/go-templ/service"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Routes(e *echo.Echo) {

	e.Validator = &CustomValidator{validator: validator.New()}

	hh := handler.NewHomeHandler()

	as := service.NewAuthService()
	ah := handler.NewAuthHandler(as)
	router.SetupRoutes(e, hh, ah)
}
