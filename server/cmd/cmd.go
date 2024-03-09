package cmd

import (
	"github.com/gx/youtubeDownloader/log"
	"github.com/gx/youtubeDownloader/server"
)

func Run() {
	srv := server.Server{}
	logger := &log.Log{}
	logger.Setup()

	srv.Logger = logger
	srv.Host()
}
