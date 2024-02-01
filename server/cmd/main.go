package main

import (
	"github.com/gx/youtubeDownloader/database"
	"github.com/gx/youtubeDownloader/internal"
	"github.com/gx/youtubeDownloader/log"
)

func main() {
	srv := internal.Server{}
	db := &database.DB{}
	logger := &log.Log{}

	logger.Setup()
	logger.Info("Starting internal")

	srv.Logger = logger
	srv.Database = db

	srv.Host()
}
