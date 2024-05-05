package cmd

import (
	"context"
	"github.com/gx/youtubeDownloader/internal/logger"
	"github.com/gx/youtubeDownloader/internal/transporters"
)

func Run() {
	ctx := context.Background()

	logger.Init(ctx)

	t := transporters.New(ctx)
	t.InitAllServers(ctx)
}
