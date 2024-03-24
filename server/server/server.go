package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/database"
	"github.com/gx/youtubeDownloader/log"
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
	conf, err := config.Get()
	if err != nil {
		log.Fatal(err, "Failed to get configuration")
	}

	s.Config = conf

	log.Info("Connecting to the database")
	db, err := database.Connect(s.Config.DatabaseEnv)
	if err != nil {
		log.Fatal(err, "Database connection failed")
	}
	s.Database = db

	log.Info("Creating router")
	gin.SetMode(gin.ReleaseMode)

	s.Router = gin.Default()

	log.Info("Setting routes")
	s.SetupRoutes()

	go func() {
		lis, err := net.Listen("tcp", ":17000")
		if err != nil {
			log.Fatal(err, "Failed to start GRPC server")
		}
		grpcServer := grpc.NewServer()
		pb.RegisterYtdlServer(grpcServer, s)
		reflection.Register(grpcServer)
		log.Info("GRPC server listening @%s", lis.Addr().String())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err, "Failed to serve GRPC server")
		}
	}()

	log.Info("Running HTTP server")
	if err := s.Router.Run(":7000"); err != nil {
		log.Fatal(err, "Failed to start server")
	}
}
