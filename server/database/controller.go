package database

import (
	"github.com/google/uuid"
	c "github.com/gx/youtubeDownloader/constants"
	"time"
)

var _ controllers = (*Database)(nil)

func (db *Database) NewSession() *Session {
	now := time.Now()
	return &Session{
		Id:        uuid.NewString(),
		CreatedAt: &now,
		UpdatedAt: &now,
		CreatedBy: "admin",
		Session:   uuid.NewString(),
		LastLogin: &now,
	}
}

func (db *Database) NewYTDL(url, storePath, sessionId string, fileSize int) *Ytdl {
	now := time.Now()
	return &Ytdl{
		Id:        uuid.NewString(),
		CreatedAt: &now,
		UpdatedAt: &now,
		CreatedBy: "admin",
		Url:       url,
		StorePath: storePath,
		SessionId: sessionId,
		Ttl:       c.FileTtl,
		Active:    true,
		FileSize:  fileSize,
	}
}
