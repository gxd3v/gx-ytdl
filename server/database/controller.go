package database

import (
	"github.com/google/uuid"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/models"
	"time"
)

func (db *Database) NewSession() *models.Session {
	now := time.Now()
	return &models.Session{
		Id:        uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: "admin",
		Session:   uuid.NewString(),
		LastLogin: now,
	}
}

func (db *Database) NewYTDL(url, storePath, sessionId string, fileSize int) *models.Ytdl {
	now := time.Now()
	return &models.Ytdl{
		Id:        uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: "admin",
		Url:       url,
		StorePath: storePath,
		SessionId: sessionId,
		Ttl:       c.FileTtl,
		Active:    true,
		FileSize:  fileSize,
	}
}
