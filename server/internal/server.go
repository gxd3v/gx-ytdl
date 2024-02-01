package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gx/youtubeDownloader/config"
)

func (s *Server) Host() {
	conf := &config.Config{}
	if co, err := conf.Get(); err != nil {
		s.Logger.Error(err.Error())
	} else {
		s.Config = co
		s.Database.Config = co
		s.Logger.Info("Creating router")
		gin.SetMode(gin.ReleaseMode)

		s.Router = gin.Default()

		s.Logger.Info("Setting routes")
		s.SetupRoutes()

		s.Logger.Info("Running the internal")
		if err := s.Router.Run(); err != nil {
			s.Logger.Info(fmt.Sprintf("Failed to start internal for this reason: %v", err.Error()))
			return
		}
	}

}
