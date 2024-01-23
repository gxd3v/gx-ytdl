package main

import (
	"github.com/gx/youtubeDownloader/database"
	"github.com/gx/youtubeDownloader/log"
	"github.com/gx/youtubeDownloader/server"
)

func main() {
	srv := server.Server{}
	db := &database.DB{}
	logger := &log.Log{}

	logger.Setup()
	logger.Info("Starting server")

	srv.Logger = logger
	srv.Database = db

	srv.Host()
}
