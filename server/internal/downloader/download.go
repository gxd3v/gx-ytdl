package downloader

import (
	"context"
	"github.com/google/uuid"
	"github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/internal/logger"
	"github.com/kkdai/youtube"
	"os"
	"path"
	"time"
)

func Downloader(ctx context.Context, url string) (bool, error) {
	sessionId, _ := ctx.Value("session").(string)

	log := logger.Fields(map[string]any{
		"where":   "downloader",
		"session": sessionId,
	})

	log.Info().Msg("Beginning download")

	dl := youtube.NewYoutube(true, false)

	t := time.Now()

	data := &Ytdl{
		Id:        uuid.NewString(),
		CreatedAt: t,
		UpdatedAt: t,
		CreatedBy: "admin",
		Url:       url,
		StorePath: constants.Temp,
		SessionId: sessionId,
		Ttl:       constants.FileTtl,
		Active:    true,
	}

	fileName := "download"

	if err := dl.StartDownloadWithHighQuality(data.StorePath, fileName, "hd1080"); err != nil {
		log.Err(err).Msg("failed to download youtube video")
		return false, err
	}

	info, err := os.Stat(path.Join(data.StorePath, fileName))
	if err != nil {
		logger.Err(err).Msg("failed to get downloaded file file")
		return false, err
	}

	data.FileSize = info.Size()

	//TODO store data on database

	log.Info().Msg("Download successfully finished")

	return true, nil
}
