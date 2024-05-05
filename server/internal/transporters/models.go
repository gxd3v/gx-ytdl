package transporters

import (
	"context"
	"github.com/gx/youtubeDownloader/config"
)

type Transporters struct {
	cfg *config.Config
}

type transporter interface {
	InitAllServers(ctx context.Context)

	initHTTP(ctx context.Context, cfg config.HTTP)
	initRPC(ctx context.Context, cfg config.RPC)
	initWS(ctx context.Context, cfg config.Websocket)
}
