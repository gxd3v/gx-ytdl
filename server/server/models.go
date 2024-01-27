package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/database"
	"github.com/gx/youtubeDownloader/log"
	"github.com/gx/youtubeDownloader/protos"
)

type Business interface {
	UpgradeConnection(ctx *gin.Context)
	Download(ctx *gin.Context, request *protos.DownloadRequest) (*protos.DownloadResponse, error)
	CreateSessionFolder(ctx *gin.Context, request *protos.CreateSessionFolderRequest) (*protos.CreateSessionFolderResponse, error)
	ListFiles(ctx *gin.Context, request *protos.ListFilesRequest) (*protos.ListFilesResponse, error)
	SendFileToClient(ctx *gin.Context, request *protos.SendFileToClientRequest) (*protos.SendFileToClientResponse, error)
	DeleteFile(ctx *gin.Context, request *protos.DeleteFileRequest) (*protos.DeleteFileResponse, error)
	DeleteSession(ctx *gin.Context, request *protos.DeleteSessionRequest) (*protos.DeleteSessionResponse, error)
}

type Server struct {
	Router    *gin.Engine
	Config    *config.Config
	Logger    *log.Log
	Ws        *websocket.Conn
	SessionID string `json:"sessionID,omitempty"`
	Storage   string `json:"storage,omitempty"`
	Database  *database.DB
}

type Route struct {
	Path       string `json:"path"`
	Method     string `json:"method"`
	Controller func() `json:"-"`
	RateLimit  int32  `json:"rateLimit"`
}
