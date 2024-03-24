package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/log"
	"github.com/gx/youtubeDownloader/models"
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

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		if s.banned(r.RemoteAddr) {
			status := http.StatusUnauthorized
			ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "You are banned from using this service"))
			return false
		}
		return true
	}

	log.Info("Upgrading connection")
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		status := http.StatusInternalServerError
		ctx.JSON(status, util.ResponseJSONBody(fmt.Sprintf("%d", status), "Failed to upgrade connection to websocket"))
		return
	}

	log.Info("Connection upgraded")
	s.Ws = ws

	log.Info("Checking for connection reopening")
	s.SessionID = ctx.Param(c.SessionParameter)

	log.Info("Listening for codes")
	s.startListener(ctx)
}

func (s *Server) startListener(ctx *gin.Context) {
	defer func() {
		if msg := recover(); msg != nil {
			s.sendMessage(ctx, &pb.PanicResponse{
				Code:    pb.ErrorsEnum_CATASTROPHIC_ERROR,
				Message: fmt.Sprintf("%s\n%+v", "Failed to start listener\n", msg),
			})
		}
	}()

	session := s.newSession(ctx)
	if session == nil {
		s.sendMessage(ctx, &pb.PanicResponse{
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
				log.Warn("Connection closed on client: %v - %v", s.SessionID, err.Error())
				_ = s.Ws.Close()
				break
			}
			log.Info(fmt.Sprintf("Failed to read message from connection %v", err.Error()))
			continue
		}

		log.Info("Got a message", base64.StdEncoding.EncodeToString(message))

		messageActionCode := pb.ActionCode{}

		err = json.Unmarshal(message, &messageActionCode)
		if ok := s.checkMessageError(ctx, err); !ok {
			continue
		}

		log.Info("Checking which action to take")

		code := pb.ActionsEnum_value[messageActionCode.Code]

		switch code {
		case int32(pb.ActionsEnum_DOWNLOAD_AUDIO):
			log.Info("Starting audio download")
			request := &pb.DownloadRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				download, _ := s.Download(ctx, request)
				s.sendMessage(ctx, download)
				continue
			}

		case int32(pb.ActionsEnum_LIST_FILES):
			log.Info("Listing files")
			files, _ := s.ListFiles(ctx, &emptypb.Empty{})
			s.sendMessage(ctx, files)
			continue

		case int32(pb.ActionsEnum_DELETE_FILE):
			log.Info("Deleting a file")
			request := &pb.DeleteFileRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				files, _ := s.DeleteFile(ctx, request)
				s.sendMessage(ctx, files)
				continue
			}

		case int32(pb.ActionsEnum_RETRIEVE_FILE):
			log.Info("Sending a file to a client")
			request := &pb.SendFileToClientRequest{}

			err = json.Unmarshal(message, &request)
			if ok := s.checkMessageError(ctx, err); ok {
				file, _ := s.SendFileToClient(ctx, request)
				s.sendMessage(ctx, file)
				continue
			}

		default:
			log.Info("Message didn't have a known code")
			s.sendMessage(ctx, &pb.PanicResponse{
				Code:    pb.ErrorsEnum_NOT_RECOGNIZED,
				Message: fmt.Sprintf("The code %v sent was not recognized", messageActionCode.Code),
			})
			continue
		}

	}
}

func (s *Server) newSession(ctx *gin.Context) *models.Session {
	database := s.Database.Transactional()
	defer func() { database.Commit() }()

	session := &models.Session{}
	if s.SessionID != "" {
		scan, err := database.GetByField(session, models.Condition{Field: "session", Operator: "=", Value: s.SessionID})
		if err != nil {
			session = s.Database.NewSession()
		}

		data, ok := scan.(*models.Session)
		if !ok {
			session = s.Database.NewSession()
			_, err := database.Insert(session)
			if err != nil {
				database.Rollback()
				return nil
			}
		} else {
			now := time.Now()
			session = data
			session.UpdatedAt = now
			session.LastLogin = now

			_, err := database.Update(session)
			if err != nil {
				database.Rollback()
				return nil
			}
		}
	} else {
		session = s.Database.NewSession()
		s.SessionID = session.Session
		log.Info("Creating a new session", session.Session)

		_, err := database.Insert(session)
		if err != nil {
			database.Rollback()
			log.Error(err, "Failed to create session")
			return nil
		}
	}

	s.sendMessage(ctx, &pb.CreateSessionResponse{
		Code:      pb.SuccessEnum_SESSION_ID,
		SessionId: session.Session,
	})

	//log.SetSessionID(session.Session)
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
		log.Error(err, "Message was malformed")
		s.sendMessage(ctx, &pb.PanicResponse{
			Code:    pb.ErrorsEnum_MALFORMED_MESSAGE,
			Message: "Message was malformed",
		})
		return false
	}

	return true
}

func (s *Server) sendMessage(_ context.Context, message proto.Message) {
	marshaller := protojson.MarshalOptions{
		EmitDefaultValues: true,
	}

	out, err := marshaller.Marshal(message)
	if err != nil {
	}

	if err = s.Ws.WriteMessage(websocket.TextMessage, out); err != nil {
		log.Error(err, "Failed to send message to client")
	}
}

func (s *Server) banned(remote string) bool {
	database := s.Database.Transactional()
	defer func() { database.Commit() }()

	if strings.Contains(remote, ":") {
		remote = strings.Split(remote, ":")[0]
	}

	bannedIP := &models.BannedIP{}
	data, err := database.GetByField(bannedIP, models.Condition{Field: "ip", Operator: "=", Value: remote})
	if err != nil {
		database.Rollback()
		log.Error(err, "Failed to scan for data")
		return false
	}

	bip, ok := data.(*models.BannedIP)
	if !ok {
		database.Rollback()
		log.Error(errors.New("cast failed"), "Failed to cast data to expected model")
		return false
	}

	if bip.Ip == remote {
		database.Rollback()
		log.Warn(fmt.Sprintf("IP %s tried to connect but it's banned from the service", remote))
		return true
	}

	return false
}
