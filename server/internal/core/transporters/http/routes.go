package http

import (
	"context"
	"github.com/gobuffalo/buffalo"
)

var basePath = "/v1"

func SetupRoutes(ctx context.Context, app *buffalo.App) {
	app.Group(basePath)

	app.GET("/session", func(c buffalo.Context) error {
		ctx = context.WithValue(ctx, "buffalo", c)
		return nil
	})
}
