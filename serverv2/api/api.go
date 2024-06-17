package api

import (
	"context"
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gx/gx-ytdl/serverv2/api/ws"
	"github.com/gx/gx-ytdl/serverv2/pkg/logger"
)

func StartServer(ctx context.Context, host, port string) error {
	app := buffalo.New(buffalo.Options{
		Host:    fmt.Sprintf("%s:%s", host, port),
		Context: ctx,
	})

	defer app.Stop(nil)

	app.GET("/connect", ws.UpgradeConnection)

	if err := app.Serve(); err != nil {
		logger.Err(err).Msg("failed to start http server")
		return err
	}

	return nil
}
