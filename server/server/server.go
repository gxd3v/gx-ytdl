package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/database"
	pb "github.com/gx/youtubeDownloader/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var _ pb.YtdlServer = (*Server)(nil)

func (s *Server) Host() {
	conf := &config.Config{}
	if co, err := conf.Get(); err != nil {
		s.Logger.Error(err.Error())
		panic(err)
	} else {
		s.Config = co

		s.Database = &database.Database{}
		s.Database = s.Database.Connect(s.Config.Database)

		s.Logger.Config = s.Config
		s.Logger.Info("Connecting to the database")

		s.Logger.Info("Creating router")
		gin.SetMode(gin.ReleaseMode)

		s.Router = gin.Default()

		s.Logger.Info("Setting routes")
		s.SetupRoutes()

		go func() {
			lis, err := net.Listen("tcp", ":17000")
			if err != nil {
				s.Logger.Error("Failed to start GRPC server")
			}
			grpcServer := grpc.NewServer()
			pb.RegisterYtdlServer(grpcServer, s)
			reflection.Register(grpcServer)
			s.Logger.Info("GRPC server listening @", lis.Addr())
			if err := grpcServer.Serve(lis); err != nil {
				s.Logger.Error("Failed to serve GRPC server")
			}
		}()

		s.Logger.Info("Running the server")
		if err := s.Router.Run(":7000"); err != nil {
			s.Logger.Info(fmt.Sprintf("Failed to start server for this reason: %v", err.Error()))
			return
		}
	}

}
