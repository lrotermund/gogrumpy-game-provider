package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v0 "github.com/lrotermund/gogrumpy-game-provider/pkg/api/v0"
)

func main() {
	e := echo.New()
	apiV0 := v0.NewAPIv0()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	games := e.Group("games/")
	e.GET("alive", isAlive)
	apiV0.Register(games)

	e.Logger.Fatal(e.Start(":8080"))
}

func isAlive(c echo.Context) error {
	return c.JSON(200, struct {
		Message string `json:"message"`
	}{Message: "ok"})
}
