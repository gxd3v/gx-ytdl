package cmd

import (
	"github.com/gx/youtubeDownloader/log"
	"github.com/gx/youtubeDownloader/server"
)

func Run() {
	log.Info("Hosting server")
	server.New().Host()
}
