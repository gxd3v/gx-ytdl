package rpc

import (
	"context"
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gx/youtubeDownloader/internal/logger"
	pb "github.com/gx/youtubeDownloader/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var _ pb.YtdlServer = (*downloader)(nil)

func StartServer(ctx context.Context, host, port string) error {
	app := buffalo.New(buffalo.Options{
		Host:    fmt.Sprintf("%s:%s", host, port),
		Context: ctx,
	})

	defer func() {
		_ = app.Stop(nil)
	}()

	for _, route := range app.Routes() {
		logger.Debug().Msgf("HTTP routes: %s", route.String())
	}

	//SetupRoutes(ctx, app)

	go func() {
		lis, err := net.Listen("tcp", ":17000")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to start GRPC internal")
			os.Exit(-1)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterYtdlServer(grpcServer, Downloader{})
		reflection.Register(grpcServer)

		logger.Info().Msgf("GRPC internal listening @%s", lis.Addr().String())

		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal().Err(err).Msg("Failed to serve GRPC internal")
			os.Exit(-1)
		}
	}()

	if err := app.Serve(); err != nil {
		logger.Err(err).Msg("failed to start http server")
		return err
	}

	return nil
}
