package ws

import (
	"encoding/base64"
	"errors"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
	"github.com/gx/gx-ytdl/serverv2/pkg/_const"
	"github.com/gx/gx-ytdl/serverv2/pkg/logger"
	"net/http"
)

func startListener(ctx buffalo.Context) error {
	defer func() {
		if msg := recover(); msg != nil {
			logger.Fatal().Msgf("panic recovered: %+v", msg)
		}
	}()

	ws, ok := ctx.Value("ws").(*websocket.Conn)
	if !ok {
		return ctx.Error(http.StatusInternalServerError, errors.New(""))
	}

	for {
		msgType, message, err := ws.ReadMessage()
		if checkClientDisconnected(ctx, msgType, err) {
			break
		}

		

	} 

	return nil
}

func checkClientDisconnected(ctx buffalo.Context, msgType int, err error) bool {
	ws, ok := ctx.Value("ws").(*websocket.Conn)
	if !ok {
		logger.Err(errors.New("no websocket in context")).Msg("context had no websocket value defined")
		return true
	}

	if err != nil {
		if msgType == _const.ClientDisconnectedCode {
			logger.Warn().Msgf("connection closed on client")
			_ = ws.Close()
			return true
		}

		logger.Err(err).Msg("failed to read client message")
		return false
	}

	return false
}
