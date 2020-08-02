package v0

import (
	"github.com/labstack/echo/v4"
)

// V0 gogrumpy-game-provider API in version v0
type V0 struct {
}

// Register registers all routes of the v0 API
func (a *V0) Register(g *echo.Group) error {
	gameGroup := g.Group("games")
	a.registerGames(gameGroup)

	return nil
}
