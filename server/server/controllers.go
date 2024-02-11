package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/database"
	"github.com/gx/youtubeDownloader/protos"
	"github.com/gx/youtubeDownloader/util"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
)

func (s *Server) UpgradeConnection(ctx *gin.Context) {
	if s.Banned(ctx) {
		status := http.StatusUnauthorized
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "You are banned from using this service"))
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	s.Logger.Info("Upgrading connection")
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "Failed to upgrade connection to websocket"))
	}

	s.Logger.Info("Connection upgraded")
	s.Ws = ws

	s.Logger.Info("Checking for connection reopening")
	s.SessionID = ctx.Param(c.SessionParameter)

	s.Logger.Info("Listening for codes")
	s.StartListener(ctx)
}

func (s *Server) StartListener(ctx *gin.Context) {
	defer func() {
		if msg := recover(); msg != nil {
			s.SendMessage(ctx, &protos.PanicResponse{
				Code:    protos.ErrorsEnum_MALFORMED_MESSAGE,
				Message: fmt.Sprintf("%s\n%+v", "Message was malformed", msg),
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
	s.CreateSessionFolder(ctx, &protos.CreateSessionFolderRequest{
		Code: protos.ActionsEnum_NEW_SESSION.String(),
		Payload: &protos.CreateSessionFolderRequestPayload{
			Session: s.SessionID,
		},
	})

	for {
		msgType, message, err := s.Ws.ReadMessage()
		if err != nil {
			if msgType == c.ClientDisconnected {
				s.Logger.Warning(fmt.Sprintf("Connection closed on client: %v - %v", s.SessionID, err.Error()))
				_ = s.Ws.Close()
				break
			}
			s.Logger.Info(fmt.Sprintf("Failed to read message from connection %v", err.Error()))
			continue
		}

		s.Logger.Info("Got a message", base64.StdEncoding.EncodeToString(message))

		messageActionCode := protos.ActionCode{}

		err = json.Unmarshal(message, &messageActionCode)
		if ok := s.checkMessageError(ctx, err); !ok {
			continue
		}

		s.Logger.Info("Checking which action to take")

		code := protos.ActionsEnum_value[messageActionCode.Code]

		switch code {
		case int32(protos.ActionsEnum_DOWNLOAD_AUDIO):
			s.Logger.Info("Starting audio download")
			request := &protos.DownloadRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				download := s.Download(ctx, request)

				s.SendMessage(ctx, download)
				continue
			}

		case int32(protos.ActionsEnum_LIST_FILES):
			s.Logger.Info("Listing files")
			files := s.ListFiles(ctx)
			s.SendMessage(ctx, files)
			continue

		case int32(protos.ActionsEnum_DELETE_FILE):
			s.Logger.Info("Deleting a file")
			request := &protos.DeleteFileRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				files := s.DeleteFile(ctx, request)
				s.SendMessage(ctx, files)
				continue
			}

		case int32(protos.ActionsEnum_RETRIEVE_FILE):
			s.Logger.Info("Sending a file to a client")
			request := &protos.SendFileToClientRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				file := s.SendFileToClient(ctx, request)

				s.SendMessage(ctx, file)
				continue
			}

		default:
			s.Logger.Info("Message didn't have a known code")
			s.SendMessage(ctx, &protos.PanicResponse{
				Code:    protos.ErrorsEnum_NOT_RECOGNIZED,
				Message: fmt.Sprintf("The code %v sent was not recognized", messageActionCode.Code),
			})
			continue
		}

	}
}

func (s *Server) checkMessageError(ctx *gin.Context, err error) bool {
	if err != nil {
		s.Logger.Error("Message was malformed", err.Error())
		s.SendMessage(ctx, &protos.PanicResponse{
			Code:    protos.ErrorsEnum_MALFORMED_MESSAGE,
			Message: "Message was malformed",
		})
		return false
	}

	return true
}

func (s *Server) SendMessage(ctx *gin.Context, message proto.Message) {
	marshaller := protojson.MarshalOptions{
		EmitDefaultValues: true,
	}

	out, err := marshaller.Marshal(message)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "The response message from the server failed to be parsed"))
	}

	if err = s.Ws.WriteMessage(websocket.TextMessage, out); err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "The response message from the server failed to be parsed"))
	}
}

func (s *Server) Banned(ctx *gin.Context) bool {
	remote := ctx.Request.RemoteAddr
	if strings.Contains(remote, ":") {
		remote = strings.Split(remote, ":")[0]
	}
	bannedIP := &database.BannedIP{}
	err := s.Database.GetByField("ip", remote).Main.Scan(bannedIP).Error
	if err != nil || bannedIP == nil {
		return false
	}

	return true
}
