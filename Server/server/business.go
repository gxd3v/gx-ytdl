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
	"os"
	"os/exec"
	"strings"
)

type Resource struct {
	Ws        *websocket.Conn
	SessionID string `json:"sessionID"`
	Storage   string `json:"storage"`
}

func (r *Resource) Download(ctx *gin.Context, audio bool, url string) {
	log.Print(fmt.Sprintf("Creating folder %s to store downloads", r.SessionID))
	r.Storage = fmt.Sprintf(c.OUTPUT_PATH, r.SessionID)
	_ = os.Mkdir(r.Storage, os.ModeAppend)

	cmd := exec.Command(c.PYTHON_BINARY, c.DOWNLOADER_PATH)
	cmd.Args = append(cmd.Args, "-u", url)
	cmd.Args = append(cmd.Args, "-op", fmt.Sprintf(c.OUTPUT_PATH, r.SessionID))
	if audio {
		cmd.Args = append(cmd.Args, "-a")
	}

	err := cmd.Run()
	if err != nil {
		r.SendMessage(ctx, c.CODE_ERROR_DOWNLOAD_FAILED, err.Error())
		return
	}

	go func() {
		_ = cmd.Wait()
		r.SendMessage(ctx, c.CODE_SUCCESS_VIDEO_DOWNLOADABLE, "path.join")
	}()
}

func (r *Resource) SendMessage(ctx *gin.Context, code, message string) {
	success := true
	if strings.Contains(code, c.CODE_ERROR_LETTER) {
		log.Print(fmt.Sprintf("ERROR: %s", message))
		success = false
	}

	resp := &models.WebsocketServerResponse{
		Id:      uuid.New().String(),
		Success: success,
		Data: models.JSONBodyMessage{
			Code:    code,
			Message: message,
		},
	}

	out, err := json.Marshal(resp)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_SERVER_RESPONSE_FAILED))
	}

	err = r.Ws.WriteMessage(websocket.TextMessage, out)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_SERVER_RESPONSE_FAILED))
	}
}
