package downloader

import (
	"context"
	"errors"
	"fmt"
	"github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/internal/core/logger"
	"github.com/gx/youtubeDownloader/internal/orchestrator/downloader/helpers"
	"os"
	"time"
)

func (d *Downloader) ListFiles(ctx context.Context) ([]*File, error) {
	sessionId, err := helpers.GetContextSession(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to list files")
		return nil, err
	}

	//TODO get files from DB and calculate TTL
	_ = sessionId
	sessionFiles := make([]*Ytdl, 0)

	files := make([]*File, 0)

	for _, sf := range sessionFiles {
		exists, err := helpers.CheckPathExists(sf.StorePath)
		if err != nil || !exists {
			logger.Err(err).Msg("file does not exist")
			continue
		}

		stat, err := os.Stat(sf.StorePath)
		if err != nil {
			logger.Err(err).Msg("something went wrong listing file")
			continue
		}

		duration, _ := time.ParseDuration(fmt.Sprintf("%d", constants.FileTtl))

		files = append(files, &File{
			Name: stat.Name(),
			Size: stat.Size(),
			Ttl:  stat.ModTime().Add(duration).Sub(time.Now()),
		})
	}

	return files, nil
}

func (d *Downloader) RetrieveFile(ctx context.Context) (bool, error) {
	return false, nil
}

func (d *Downloader) DeleteFile(ctx context.Context, fileId string) (bool, error) {
	sessionId, err := helpers.GetContextSession(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to delete file")
		return false, err
	}

	//TODO get file from database filtering by id and session
	_ = sessionId
	data := new(Ytdl)

	exists, err := helpers.CheckPathExists(data.StorePath)
	if err != nil || !exists {
		logger.Err(err).Msg("file does not exist")
		return false, errors.New(err.Error()) //this forces a new error even if err is nil but exists is false
	}

	if err := os.Remove(data.StorePath); err != nil {
		logger.Err(err).Msg("failed to delete file")
		return false, err
	}

	t := time.Now()

	data.DeletedAt = &t
	data.UpdatedAt = t
	data.Active = false
	//TODO update DB with new deleted time

	return true, nil
}
