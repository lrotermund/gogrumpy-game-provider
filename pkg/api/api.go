package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v0 "github.com/lrotermund/gogrumpy-game-provider/pkg/api/v0"
	"github.com/lrotermund/gogrumpy-game-provider/pkg/config"
)

// API is the base api interface
type API interface {
	Register(*echo.Group) error
}

// NewV0 returns the implementation of the v0 api.
func NewV0() API {
	return &v0.V0{}
}

// New sets up the complete API and returns it. The API can be executed as it is returned.
func New(config config.Configuration) (*echo.Echo, error) {
	initHealthCheck()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	v0Group := e.Group("api/v0/")

	v0 := NewV0()
	if err := v0.Register(v0Group); err != nil {
		panic(err)
	}

	return e, nil
}
