package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/log"
	"github.com/gx/youtubeDownloader/models"
	"github.com/gx/youtubeDownloader/util"
	"net/http"
	"strings"
)

func (r *Resource) UpgradeConnection(ctx *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	log.Print("Upgrading connection")
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_UPGRADE_FAILED))
	}

	log.Print("Connection upgraded")
	r.Ws = ws

	log.Print("Listening for codes")
	r.StartListener(ctx)
}

func (r *Resource) StartListener(ctx *gin.Context) {
	defer func() {
		if msg := recover(); msg != nil {
			r.SendMessage(ctx, c.CODE_ERROR_MALFORMED_MESSAGE, fmt.Sprintf("%s\n%+v", c.TEXT_ERROR_MALFORMED_MESSAGE, msg))
			r.StartListener(ctx)
		}
	}()

	if r.SessionID == "" {
		r.SessionID = uuid.New().String()
	}

	for {
		_, message, err := r.Ws.ReadMessage()
		if err != nil {
			log.Print(fmt.Sprintf("Failed to read message from connection %v", err.Error()))
			continue
		}
		log.Print("Got a message")
		log.Print(fmt.Sprintf("%+v", string(message)))

		msg := models.WebsocketMessage{}

		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Print("Message was malformed")
			r.SendMessage(ctx, c.CODE_ERROR_MALFORMED_MESSAGE, c.TEXT_ERROR_MALFORMED_MESSAGE)
			continue
		}

		log.Print("Parsed message")
		log.Print(fmt.Sprintf("%+v", msg))

		log.Print("Checking which action to take")
		switch strings.ToUpper(msg.Code) {
		case c.CODE_DOWNLOAD_AUDIO:
			r.Download(ctx, true, msg.Payload["url"].(string))
			continue

		case c.CODE_DOWNLOAD_VIDEO_AUDIO:
			r.Download(ctx, false, msg.Payload["url"].(string))
			continue

		default:
			r.SendMessage(ctx, c.CODE_ERROR_CODE_NOT_RECOGNIZED, fmt.Sprintf(c.TEXT_ERROR_CODE_NOT_RECOGNIZED, msg.Code))
			continue
		}

	}
}
