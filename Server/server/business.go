package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/models"
	"github.com/gx/youtubeDownloader/util"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func (s *Server) Download(ctx *gin.Context, audio bool, url string) {
	cmd := exec.Command(s.Config.PythonBinary, s.Config.DownloaderPath)
	cmd.Args = append(cmd.Args, "-u", url)
	cmd.Args = append(cmd.Args, "-op", fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if audio {
		cmd.Args = append(cmd.Args, "-a")
	}

	err := cmd.Run()
	if err != nil {
		s.SendMessage(ctx, c.CODE_ERROR_DOWNLOAD_FAILED, err.Error())
		return
	}

	go func() {
		_ = cmd.Wait()
		s.SendMessage(ctx, c.CODE_SUCCESS_VIDEO_DOWNLOADABLE, "File is done downloading")
	}()
}

func (s *Server) CreateSessionFolder() {
	s.Logger.Info(fmt.Sprintf("Creating folder %s to store downloads", s.SessionID))
	s.Storage = fmt.Sprintf(s.Config.OutputPath, s.SessionID)
	_ = os.Mkdir(s.Storage, os.ModeAppend)
}

func (s *Server) ListFiles(ctx *gin.Context) {
	if files, err := os.ReadDir(fmt.Sprintf(s.Config.OutputPath, s.SessionID)); err != nil {
		s.Logger.Error(c.TEXT_ERROR_FAILED_LISTING_FILES, err.Error())
		s.SendMessage(ctx, c.CODE_ERROR_FAILED_LISTING_FILES, c.TEXT_ERROR_FAILED_LISTING_FILES)
	} else {
		if len(files) == 0 {
			s.Logger.Warning("No files in folder to show")
			s.SendMessage(ctx, c.CODE_SUCCESS_LISTED_FILES, base64.StdEncoding.EncodeToString([]byte("{}")))
		} else {
			output := map[int]string{}

			for index, file := range files {
				output[index] = file.Name()
			}

			out, _ := json.Marshal(output)
			s.Logger.Info("files in the session", base64.StdEncoding.EncodeToString(out))
			s.SendMessage(ctx, c.CODE_SUCCESS_LISTED_FILES, base64.StdEncoding.EncodeToString(out))
		}

	}
}

func (s *Server) SendFileToClient(ctx *gin.Context, fileName string) {
	s.Logger.Info("Client requested file", fileName)
	s.SendMessage(ctx, c.CODE_SUCCESS_READY_TO_SEND, fileName)
}

func (s *Server) DeleteFile(ctx *gin.Context, fileName string) {
	err := os.Remove(path.Join(fmt.Sprintf(s.Config.OutputPath, s.SessionID), fileName))
	if err != nil {
		s.Logger.Error(c.TEXT_ERROR_FAILED_DELETE_FILE)
		s.SendMessage(ctx, c.CODE_ERROR_FAILED_DELETE_FILE, c.TEXT_ERROR_FAILED_DELETE_FILE)
	} else {
		s.Logger.Info("Client removed file", fileName)
		s.SendMessage(ctx, c.CODE_SUCCESS_DELETE_FILE, "File deleted")
		s.ListFiles(ctx)
	}
}

func (s *Server) DeleteSession(ctx *gin.Context) {
	err := os.Remove(fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if err != nil {
		s.Logger.Error(c.TEXT_ERROR_FAILED_DELETE_SESSION)
		s.SendMessage(ctx, c.CODE_SUCCESS_SESSION_DELETE, c.TEXT_ERROR_FAILED_DELETE_SESSION)
	} else {
		s.Logger.Info("Client removed session")
		s.SendMessage(ctx, c.CODE_SUCCESS_SESSION_DELETE, "Session deleted")

		_ = s.Ws.Close()
		s.SessionID = ""
	}
}

func (s *Server) SendMessage(ctx *gin.Context, code, message string) {
	success := true
	if strings.Contains(code, c.CODE_ERROR_LETTER) {
		s.Logger.Error(fmt.Sprintf("ERROR: %s", message))
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

	err = s.Ws.WriteMessage(websocket.TextMessage, out)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_SERVER_RESPONSE_FAILED))
	}
}
