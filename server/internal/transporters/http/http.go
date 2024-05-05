package http

import (
	"context"
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gx/youtubeDownloader/internal/logger"
)

func StartServer(ctx context.Context, host, port string) error {
	app := buffalo.New(buffalo.Options{
		Host:    fmt.Sprintf("%s:%s", host, port),
		Context: ctx,
	})

	defer func() {
		_ = app.Stop(nil)
	}()

	for _, route := range app.Routes() {
		logger.Debug().Msgf("HTTP routes: %s", route.String())
	}

	SetupRoutes(ctx, app)

	if err := app.Serve(); err != nil {
		logger.Err(err).Msg("failed to start http server")
		return err
	}

	return nil
}
