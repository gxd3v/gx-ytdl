package datastore

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gx/youtubeDownloader/database"
	"github.com/gx/youtubeDownloader/models"
)

func NewDownloaderRepo(db *pop.Connection) database.Generic[models.Ytdl] {
	return database.New[models.Ytdl](db)
}
