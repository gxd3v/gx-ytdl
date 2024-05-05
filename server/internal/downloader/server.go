package downloader

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/database"
	logger2 "github.com/gx/youtubeDownloader/internal/logger"
	"github.com/gx/youtubeDownloader/logger"
	pb "github.com/gx/youtubeDownloader/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var _ pb.YtdlServer = (*Server)(nil)

func New() *Server {
	return &Server{}
}

func (s *Server) Host() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(errors.New("panic recovery"), "panicked: %+v", err)
		}
	}()

	s.Config = config.Get()

	logger2.Info("Connecting to the database")
	db, err := database.Connect(s.Config.DatabaseEnv)
	if err != nil {
		logger2.Fatal(err, "Database connection failed")
	}
	s.Database = db

	logger2.Info("Creating router")
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	logger2.Info("Setting routes")
	s.SetupRoutes(router)

	go func() {
		lis, err := net.Listen("tcp", ":17000")
		if err != nil {
			logger2.Fatal(err, "Failed to start GRPC internal")
		}
		grpcServer := grpc.NewServer()
		pb.RegisterYtdlServer(grpcServer, s)
		reflection.Register(grpcServer)
		logger2.Info("GRPC internal listening @%s", lis.Addr().String())
		if err := grpcServer.Serve(lis); err != nil {
			logger2.Fatal(err, "Failed to serve GRPC internal")
		}
	}()

	logger2.Info("Running HTTP internal")
	if err := router.Run(":7000"); err != nil {
		logger2.Fatal(err, "Failed to start internal")
	}
}
