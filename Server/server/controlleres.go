package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/models"
	"github.com/gx/youtubeDownloader/util"
	"net/http"
	"strings"
)

func (s *Server) UpgradeConnection(ctx *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	s.Logger.Info("Upgrading connection")
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_UPGRADE_FAILED))
	}

	s.Logger.Info("Connection upgraded")
	s.Ws = ws

	s.Logger.Info("Listening for codes")
	s.StartListener(ctx)
}

func (s *Server) StartListener(ctx *gin.Context) {
	defer func() {
		if msg := recover(); msg != nil {
			s.SendMessage(ctx, c.CODE_ERROR_MALFORMED_MESSAGE, fmt.Sprintf("%s\n%+v", c.TEXT_ERROR_MALFORMED_MESSAGE, msg))
			s.StartListener(ctx)
		}
	}()

	if s.SessionID == "" {
		s.SessionID = uuid.New().String()
	}

	for {
		_, message, err := s.Ws.ReadMessage()
		if err != nil {
			s.Logger.Info(fmt.Sprintf("Failed to read message from connection %v", err.Error()))
			continue
		}
		s.Logger.Info("Got a message")
		s.Logger.Info(fmt.Sprintf("%+v", string(message)))

		msg := models.WebsocketMessage{}

		err = json.Unmarshal(message, &msg)
		if err != nil {
			s.Logger.Info("Message was malformed")
			s.SendMessage(ctx, c.CODE_ERROR_MALFORMED_MESSAGE, c.TEXT_ERROR_MALFORMED_MESSAGE)
			continue
		}

		s.Logger.Info("Parsed message")
		s.Logger.Info(fmt.Sprintf("%+v", msg))

		s.Logger.Info("Checking which action to take")
		switch strings.ToUpper(msg.Code) {
		case c.CODE_DOWNLOAD_AUDIO:
			s.Download(ctx, true, msg.Payload["url"].(string))
			continue

		case c.CODE_DOWNLOAD_VIDEO_AUDIO:
			s.Download(ctx, false, msg.Payload["url"].(string))
			continue

		default:
			s.SendMessage(ctx, c.CODE_ERROR_CODE_NOT_RECOGNIZED, fmt.Sprintf(c.TEXT_ERROR_CODE_NOT_RECOGNIZED, msg.Code))
			continue
		}

	}
}
