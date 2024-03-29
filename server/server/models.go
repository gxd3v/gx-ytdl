package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/database"
)

//type Business interface {
//	UpgradeConnection(ctx *gin.Context)
//	Download(ctx *gin.Context, request *protos.DownloadRequest) *protos.DownloadResponse
//	CreateSessionFolder(ctx *gin.Context, request *protos.CreateSessionFolderRequest) *protos.CreateSessionFolderResponse
//	ListFiles(ctx *gin.Context) *protos.ListFilesResponse
//	SendFileToClient(ctx *gin.Context, request *protos.SendFileToClientRequest) *protos.SendFileToClientResponse
//	DeleteFile(ctx *gin.Context, request *protos.DeleteFileRequest) *protos.DeleteFileResponse
//	DeleteSession(ctx *gin.Context) *protos.DeleteSessionResponse
//}

type Server struct {
	Router    *gin.Engine
	Config    *config.Config
	Ws        *websocket.Conn
	SessionID string `json:"sessionID,omitempty"`
	Storage   string `json:"storage,omitempty"`
	Database  *database.Database
}

type Route struct {
	Path       string `json:"path"`
	Method     string `json:"method"`
	Controller func() `json:"-"`
	RateLimit  int32  `json:"rateLimit"`
}
