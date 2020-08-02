package v0

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lrotermund/gogrumpy-game-provider/pkg/game"
)

func (a *V0) registerGames(g *echo.Group) {
	g.GET("", a.gamesHandler)
}

func (a *V0) gamesHandler(c echo.Context) error {
	boardgame := game.Game{
		ID: uuid.New(),
		PreviewImage: game.Image{
			Src:  `/demo/robowars.png`,
			Name: `robowars.png`,
			Mime: `image/png`,
		},
		Name:        `RoboWars`,
		Description: `In einem post-apokalyptischen Zeitalter kämpfen verbliebene Clans um die Rohstoffvorherrschaft.`,
		LongDescription: `
		<p>
			Die Welt ist am Ende, es gibt keine Ressourcen mehr und um das Wenige verbliebene müssen 
			die bitter verfeindeten Clans untereinander kämpfen.
	  	</p>
	  	<p class="mt-2">
			Programmiere deinen Robo mit dem richtigen Go-Code um deine Gegner auszuschalten und um an die 
			verbliebenen Ressourcen zu gelangen.
	  	</p>`,
	}

	return c.JSON(200, []game.Game{boardgame})
}
