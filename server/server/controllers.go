package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/protos"
	"github.com/gx/youtubeDownloader/util"
	"net/http"
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

	s.Logger.Info("Checking for connection reopening")
	s.SessionID = ctx.Param(c.SESSION_PARAMETER)

	s.Logger.Info("Listening for codes")
	s.StartListener(ctx)
}

func (s *Server) StartListener(ctx *gin.Context) {
	defer func() {
		if msg := recover(); msg != nil {
			s.SendMessage(ctx, &protos.PanicResponse{
				Code:    protos.ErrorsEnum_MALFORMED_MESSAGE,
				Message: fmt.Sprintf("%s\n%+v", c.TEXT_ERROR_MALFORMED_MESSAGE, msg),
			})
			s.StartListener(ctx)
		}
	}()

	if s.SessionID == "" {
		s.SessionID = uuid.New().String()
		s.Logger.Info("Creating a new session", s.SessionID)
		s.SendMessage(ctx, &protos.CreateSessionResponse{
			Code:      protos.SuccessEnum_SESSION_ID,
			SessionId: s.SessionID,
		})
	}

	s.Logger.SetSessionID(s.SessionID)
	_, err := s.CreateSessionFolder(ctx, &protos.CreateSessionFolderRequest{
		Code: protos.ActionsEnum_NEW_SESSION.String(),
		Payload: &protos.CreateSessionFolderRequestPayload{
			Session: s.SessionID,
		},
	})
	if err != nil {
		s.SendMessage(ctx, &protos.PanicResponse{
			Code:    protos.ErrorsEnum_FOLDER_ALREADY_EXISTS,
			Message: err.Error(),
		})
	}

	for {
		msgType, message, err := s.Ws.ReadMessage()
		if err != nil {
			if msgType == c.CLIENT_DISCONNECTED {
				s.Logger.Warning(fmt.Sprintf("Connection closed on client: %v - %v", s.SessionID, err.Error()))
				_ = s.Ws.Close()
				break
			}
			s.Logger.Info(fmt.Sprintf("Failed to read message from connection %v", err.Error()))
			continue
		}

		s.Logger.Info("Got a message", base64.StdEncoding.EncodeToString(message))

		messageActionCode := struct {
			Code string `json:"code"`
		}{}

		err = json.Unmarshal(message, &messageActionCode)
		if ok := s.checkMessageError(ctx, err); !ok {
			continue
		}

		s.Logger.Info("Checking which action to take")

		code := protos.ActionsEnum_value[messageActionCode.Code]

		switch code {
		case int32(protos.ActionsEnum_DOWNLOAD_AUDIO):
			request := &protos.DownloadRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				download, _ := s.Download(ctx, request)

				s.SendMessage(ctx, download)
				continue
			}

		//case c.CODE_DOWNLOAD_VIDEO_AUDIO:
		//	s.Download(ctx, false, msg.Payload["url"].(string))
		//	continue
		//
		//case c.CODE_LIST_FILES:
		//	s.ListFiles(ctx)
		//	continue
		//
		//case c.CODE_SEND_FILE_TO_CLIENT:
		//	s.SendFileToClient(ctx, msg.Payload["name"].(string))
		//	continue
		//
		//case c.CODE_DELETE_FILE:
		//	s.DeleteFile(ctx, msg.Payload["name"].(string))
		//	continue
		//
		//case c.CODE_DELETE_SESSION:
		//	s.DeleteSession(ctx)
		//	continue

		default:
			s.SendMessage(ctx, &protos.PanicResponse{
				Code:    protos.ErrorsEnum_NOT_RECOGNIZED,
				Message: fmt.Sprintf(c.TEXT_ERROR_CODE_NOT_RECOGNIZED, messageActionCode.Code),
			})
			continue
		}

	}
}

func (s *Server) checkMessageError(ctx *gin.Context, err error) bool {
	if err != nil {
		s.Logger.Info("Message was malformed", err.Error())
		s.SendMessage(ctx, &protos.PanicResponse{
			Code:    protos.ErrorsEnum_MALFORMED_MESSAGE,
			Message: c.TEXT_ERROR_MALFORMED_MESSAGE,
		})
		return false
	}

	return true
}

func (s *Server) SendMessage(ctx *gin.Context, message interface{}) {
	out, err := json.Marshal(message)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_SERVER_RESPONSE_FAILED))
	}

	if err = s.Ws.WriteMessage(websocket.TextMessage, out); err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_SERVER_RESPONSE_FAILED))
	}
}

//func (s *Server) SendMessage(ctx *gin.Context, code, message string) {
//	success := true
//	if strings.Contains(code, c.CODE_ERROR_LETTER) {
//		s.Logger.Error(fmt.Sprintf("ERROR: %s", message))
//		success = false
//	}
//
//	resp := &models.WebsocketServerResponse{
//		Id:      uuid.New().String(),
//		Success: success,
//		Data: models.JSONBodyMessage{
//			Code:    code,
//			Message: message,
//		},
//	}
//
//	out, err := json.Marshal(resp)
//	if err != nil {
//		status := http.StatusInternalServerError
//		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_SERVER_RESPONSE_FAILED))
//	}
//
//	err = s.Ws.WriteMessage(websocket.TextMessage, out)
//	if err != nil {
//		status := http.StatusInternalServerError
//		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), c.TEXT_ERROR_SERVER_RESPONSE_FAILED))
//	}
//}
