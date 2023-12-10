package main

import (
	"github.com/gx/youtubeDownloader/log"
	"github.com/gx/youtubeDownloader/server"
)

func main() {
	log.Print("Starting server")
	srv := server.Server{}
	srv.Host()
}
