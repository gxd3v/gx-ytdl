package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/database"
	"github.com/gx/youtubeDownloader/log"
)

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
