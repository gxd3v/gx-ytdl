package ws

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/internal/core/logger"
	pb "github.com/gx/youtubeDownloader/protos"
	"net/http"
)

func UpgradeConnection(ctx buffalo.Context) error {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		if banned(r.RemoteAddr) {
			return false
		}
		return true
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

func startListener(ctx buffalo.Context) error {
	defer func() {
		if msg := recover(); msg != nil {
			sendMessage(ctx, &pb.PanicResponse{
				Code:    pb.ErrorsEnum_CATASTROPHIC_ERROR,
				Message: fmt.Sprintf("%s\n%+v", "Failed to start listener\n", msg),
			})
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

		logger.Info().Msgf("got a message: %s", base64.StdEncoding.EncodeToString(message))

		messageActionCode := pb.ActionCode{}

		err = json.Unmarshal(message, &messageActionCode)
		if ok := checkMessageError(ctx, err); !ok {
			continue
		}

		code := pb.ActionsEnum_value[messageActionCode.Code]

		if err := checkAction(ctx, code, message); err != nil {
			return err
		}
	}

	return nil
}

func banned(remote string) bool {
	////database := s.Database.Transactional()
	////defer func() { database.Commit() }()
	//
	//if strings.Contains(remote, ":") {
	//	remote = strings.Split(remote, ":")[0]
	//}
	//
	//bannedIP := make([]*models.BannedIP, 0)
	//data, err := s.Database.GetByField(bannedIP, models.Condition{Field: "ip", Operator: "=", Value: remote})
	//if err != nil {
	//	//database.Rollback()
	//	logger.Error(err, "Failed to scan for data")
	//	return false
	//}
	//
	//bip, ok := data.(*models.BannedIP)
	//if !ok {
	//	//database.Rollback()
	//	logger.Error(errors.New("cast failed"), "Failed to cast data to expected model")
	//	return false
	//}
	//
	//if bip.Ip == remote {
	//	//database.Rollback()
	//	logger.Warn().Msg(fmt.Sprintf("IP %s tried to connect but it's banned from the service", remote))
	//	return true
	//}

	return false
}

func checkClientDisconnected(ctx buffalo.Context, msgType int, err error) bool {
	ws, ok := ctx.Value("ws").(*websocket.Conn)
	if !ok {
		logger.Err(errors.New("no websocket in context")).Msg("context had no websocket value defined")
		return true
	}

	if err != nil {
		if msgType == c.ClientDisconnected {
			logger.Warn().Msgf("connection closed on client")
			_ = ws.Close()
			return true
		}
		logger.Err(err).Msg("failed to read client message")
		return false
	}

	return false
}
