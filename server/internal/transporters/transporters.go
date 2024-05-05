package transporters

import (
	"context"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/internal/logger"
	"github.com/gx/youtubeDownloader/internal/transporters/http"
	"github.com/gx/youtubeDownloader/internal/transporters/rpc"
	"github.com/gx/youtubeDownloader/internal/transporters/ws"
	"github.com/gx/youtubeDownloader/util"
	"os"
	"strconv"
)

var _ transporter = (*Transporters)(nil)

func New(ctx context.Context) *Transporters {
	cfg := config.New(ctx)

	return &Transporters{
		cfg: cfg,
	}
}

func (t *Transporters) InitAllServers(ctx context.Context) {
	logger.Info().Msg("starting http server")
	go t.initHTTP(ctx, t.cfg.Transporters.HTTP)

	logger.Info().Msg("starting rpc server")
	go t.initRPC(ctx, t.cfg.Transporters.RPC)

	logger.Info().Msg("starting websocket server")
	go t.initWS(ctx, t.cfg.Transporters.WS)
}

func (t *Transporters) initHTTP(ctx context.Context, cfg config.HTTP) {
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start http server")
		os.Exit(-1)
	}

	for {
		if !util.CheckPortBusy(cfg.Host, strconv.Itoa(port)) {
			break
		}

		port++
		logger.Warn().Msgf("retrying with port %d", port)
	}

	logger.Info().Msgf("HTTP server is starting on port %d", port)
	if err := http.StartServer(ctx, cfg.Host, strconv.Itoa(port)); err != nil {
		logger.Fatal().Err(err).Msg("failed to start http server")
		os.Exit(-1)
	}
}

func (t *Transporters) initRPC(ctx context.Context, cfg config.RPC) {
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start http server")
		os.Exit(-1)
	}

	for {
		if !util.CheckPortBusy(cfg.Host, strconv.Itoa(port)) {
			break
		}

		port++
		logger.Warn().Msgf("retrying with port %d", port)
	}

	logger.Info().Msgf("websocket server is starting on port %d", port)
	if err := rpc.StartServer(ctx, cfg.Host, strconv.Itoa(port)); err != nil {
		logger.Fatal().Err(err).Msg("failed to start http server")
		os.Exit(-1)
	}
}

func (t *Transporters) initWS(ctx context.Context, cfg config.Websocket) {
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start http server")
		os.Exit(-1)
	}

	for {
		if !util.CheckPortBusy(cfg.Host, strconv.Itoa(port)) {
			break
		}

		port++
		logger.Warn().Msgf("retrying with port %d", port)
	}

	logger.Info().Msgf("websocket server is starting on port %d", port)
	if err := ws.StartServer(ctx, cfg.Host, strconv.Itoa(port)); err != nil {
		logger.Fatal().Err(err).Msg("failed to start http server")
		os.Exit(-1)
	}
}
