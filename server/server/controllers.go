package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	db "github.com/gx/youtubeDownloader/database"
	pb "github.com/gx/youtubeDownloader/protos"
	"github.com/gx/youtubeDownloader/util"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strings"
	"time"
)

func (s *Server) UpgradeConnection(ctx *gin.Context) {
	if s.Banned(ctx) {
		status := http.StatusUnauthorized
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "You are banned from using this service"))
		return
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
			s.SendMessage(ctx, &pb.PanicResponse{
				Code:    pb.ErrorsEnum_MALFORMED_MESSAGE,
				Message: fmt.Sprintf("%s\n%+v", "Message was malformed", msg),
			})
			s.StartListener(ctx)
		}
	}()

	session := s.NewSession(ctx)
	if session == nil {
		s.SendMessage(ctx, &pb.PanicResponse{
			Code:    pb.ErrorsEnum_CATASTROPHIC_ERROR,
			Message: "Failed to start a new session",
		})
		_ = s.Ws.Close()
		return
	}

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

		messageActionCode := pb.ActionCode{}

		err = json.Unmarshal(message, &messageActionCode)
		if ok := s.checkMessageError(ctx, err); !ok {
			continue
		}

		s.Logger.Info("Checking which action to take")

		code := pb.ActionsEnum_value[messageActionCode.Code]

		switch code {
		case int32(pb.ActionsEnum_DOWNLOAD_AUDIO):
			s.Logger.Info("Starting audio download")
			request := &pb.DownloadRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				download, _ := s.Download(ctx, request)

				s.SendMessage(ctx, download)
				continue
			}

		case int32(pb.ActionsEnum_LIST_FILES):
			s.Logger.Info("Listing files")
			files, _ := s.ListFiles(ctx, &emptypb.Empty{})
			s.SendMessage(ctx, files)
			continue

		case int32(pb.ActionsEnum_DELETE_FILE):
			s.Logger.Info("Deleting a file")
			request := &pb.DeleteFileRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				files, _ := s.DeleteFile(ctx, request)
				s.SendMessage(ctx, files)
				continue
			}

		case int32(pb.ActionsEnum_RETRIEVE_FILE):
			s.Logger.Info("Sending a file to a client")
			request := &pb.SendFileToClientRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				file, _ := s.SendFileToClient(ctx, request)

				s.SendMessage(ctx, file)
				continue
			}

		default:
			s.Logger.Info("Message didn't have a known code")
			s.SendMessage(ctx, &pb.PanicResponse{
				Code:    pb.ErrorsEnum_NOT_RECOGNIZED,
				Message: fmt.Sprintf("The code %v sent was not recognized", messageActionCode.Code),
			})
			continue
		}

	}
}

func (s *Server) NewSession(ctx *gin.Context) *db.Session {
	transaction := s.Database.Transactional()
	defer func() { _ = transaction.Commit() }()

	session := &db.Session{}
	if s.SessionID != "" {
		scan, err := s.Database.GetByField("session", s.SessionID).Scan(&session)
		if err != nil {
			session = s.Database.NewSession()
		}

		data, ok := scan.(*db.Session)
		if !ok {
			session = s.Database.NewSession()
		} else {
			now := time.Now()
			session = data
			session.UpdatedAt = &now
			session.LastLogin = &now
		}
	} else {
		session = s.Database.NewSession()
		s.SessionID = session.Session
		s.Logger.Info("Creating a new session", session.Session)
	}

	err := transaction.Insert(session)
	if err != nil {
		_ = transaction.Rollback()
		return nil
	}

	s.SendMessage(ctx, &pb.CreateSessionResponse{
		Code:      pb.SuccessEnum_SESSION_ID,
		SessionId: session.Session,
	})

	s.Logger.SetSessionID(session.Session)
	_, _ = s.CreateSessionFolder(ctx, &pb.CreateSessionFolderRequest{
		Code: pb.ActionsEnum_NEW_SESSION.String(),
		Payload: &pb.CreateSessionFolderRequestPayload{
			Session: session.Session,
		},
	})

	return session
}

func (s *Server) checkMessageError(ctx *gin.Context, err error) bool {
	if err != nil {
		s.Logger.Error("Message was malformed", err.Error())
		s.SendMessage(ctx, &pb.PanicResponse{
			Code:    pb.ErrorsEnum_MALFORMED_MESSAGE,
			Message: "Message was malformed",
		})
		return false
	}

	return true
}

func (s *Server) SendMessage(ctx context.Context, message proto.Message) {
	marshaller := protojson.MarshalOptions{
		EmitDefaultValues: true,
	}

	out, err := marshaller.Marshal(message)
	if err != nil {
		//status := http.StatusInternalServerError
		//ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "The response message from the server failed to be parsed"))
	}

	if err = s.Ws.WriteMessage(websocket.TextMessage, out); err != nil {
		//status := http.StatusInternalServerError
		//ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "The response message from the server failed to be parsed"))
	}
}

func (s *Server) Banned(ctx *gin.Context) bool {
	remote := ctx.Request.RemoteAddr
	if strings.Contains(remote, ":") {
		remote = strings.Split(remote, ":")[0]
	}

	bannedIP := &db.BannedIP{}
	data, err := s.Database.Model(bannedIP).GetByField("ip", remote).Scan(bannedIP)
	if err != nil {
		s.Logger.Error("Failed to scan for data", err.Error())
		return false
	}

	bip, ok := data.(*db.BannedIP)
	if !ok {
		s.Logger.Error("Failed to cast data to expected model")
		return false
	}

	if bip.Ip == remote {
		s.Logger.Warning(fmt.Sprintf("IP %s tried to connect but it's banned from the service", remote))
		return true
	}

	return false
}
