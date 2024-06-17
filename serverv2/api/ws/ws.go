package ws

import (
	"errors"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
	"github.com/gx/gx-ytdl/serverv2/pkg/logger"
	"net/http"
)

func UpgradeConnection(ctx buffalo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	logger.Info().Msg("upgrading connection")
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		logger.Err(err).Msg("failed to upgrade connection")
		return ctx.Error(http.StatusUnauthorized, errors.New("connection was not upgraded, no further details"))
	}

	ctx.Set("ws", ws)

	logger.Info().Msg("starting message listener")
	if err := startListener(ctx); err != nil {
		logger.Err(err).Msg("failed to start listener")
		return err
	}

	return nil
}
