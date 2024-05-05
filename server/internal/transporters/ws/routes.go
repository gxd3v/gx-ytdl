package ws

import (
	"github.com/gobuffalo/buffalo"
)

var basePath = "/v1"

func SetupRoutes(app *buffalo.App) {
	app.Group(basePath)

	app.GET("/connect", UpgradeConnection)
}
