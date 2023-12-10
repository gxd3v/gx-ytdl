package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gx/youtubeDownloader/log"
)

type Server struct {
	Router *gin.Engine
}

func (s *Server) Host() {
	log.Print("Creating router")
	gin.SetMode(gin.ReleaseMode)
	s.Router = gin.Default()
	log.Print("Setting routes")
	s.SetupRoutes()
	log.Print("Running the server")
	err := s.Router.Run()
	if err != nil {
		log.Print(fmt.Sprintf("Failed to start server for this reason: %v", err.Error()))
		return
	}
}
