package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/database"
)

func (s *Server) Host() {
	conf := &config.Config{}
	if co, err := conf.Get(); err != nil {
		s.Logger.Error(err.Error())
		panic(err)
	} else {
		s.Config = co

		s.Database = &database.DB{}
		s.Database = s.Database.Connect(s.Config.Database, database.TableName())

		s.Logger.Config = s.Config
		s.Logger.Info("Connecting to the database")

		s.Logger.Info("Creating router")
		gin.SetMode(gin.ReleaseMode)

		s.Router = gin.Default()

		s.Logger.Info("Setting routes")
		s.SetupRoutes()

		s.Logger.Info("Running the server")
		if err := s.Router.Run(); err != nil {
			s.Logger.Info(fmt.Sprintf("Failed to start server for this reason: %v", err.Error()))
			return
		}
	}

}
