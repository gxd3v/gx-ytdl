package cmd

import (
	"context"
	"github.com/gx/youtubeDownloader/internal/core/logger"
	"github.com/gx/youtubeDownloader/internal/core/transporters"
)

func Run() {
	ctx := context.Background()

	logger.Init(ctx)

	t := transporters.New(ctx)
	t.InitAllServers(ctx)
}
