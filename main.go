package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/randijulio13/go-templ/app"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	app.Routes(e)
	e.Start(":3000")
}
