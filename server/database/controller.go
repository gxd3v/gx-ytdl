package database

import (
	"github.com/google/uuid"
	c "github.com/gx/youtubeDownloader/constants"
	"time"
)

var _ controllers = (*Database)(nil)

func (db *Database) NewSession() *Session {
	return &Session{
		Id:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "admin",
		Session:   uuid.NewString(),
		LastLogin: time.Now(),
	}
}

func (db *Database) NewYTDL(url, storePath, sessionId string, fileSize int) *Ytdl {
	return &Ytdl{
		Id:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "admin",
		Url:       url,
		StorePath: storePath,
		SessionId: sessionId,
		Ttl:       c.FileTtl,
		Active:    true,
		FileSize:  fileSize,
	}
}
