package v0

import (
	"github.com/labstack/echo/v4"
	"github.com/lrotermund/gogrumpy-game-provider/pkg/api"
)

type apiV0 struct {
}

// NewAPIv0 returns a new api v0
func NewAPIv0() api.API {
	return &apiV0{}
}

func (a *apiV0) Register(g *echo.Group) error {
	g.GET("", a.handler)
	return nil
}

func (a *apiV0) handler(c echo.Context) error {
	return c.String(200, "Hello World!")
}
