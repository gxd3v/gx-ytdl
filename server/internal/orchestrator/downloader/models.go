package downloader

import (
	"context"
	"time"
)

type Ytdl struct {
	Id        string        `json:"id"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	DeletedAt *time.Time    `json:"deletedAt"`
	CreatedBy string        `json:"createdBy"`
	Url       string        `json:"url"`
	StorePath string        `json:"storePath"`
	SessionId string        `json:"sessionId"`
	Ttl       time.Duration `json:"ttl"`
	Active    bool          `json:"active"`
	FileSize  int64         `json:"fileSize"`
}

type Session struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	CreatedBy string     `json:"createdBy"`
	Session   string     `json:"session"`
	LastLogin time.Time  `json:"lastLogin"`
}

type File struct {
	Name string
	Size int64
	Ttl  time.Duration
}

type downloader interface {
	Download(ctx context.Context, url string) (bool, error)
	ListFiles(ctx context.Context) ([]*File, error)
	RetrieveFile(ctx context.Context) (bool, error)
	DeleteFile(ctx context.Context, fileId string) (bool, error)
	CreateOrGetSession(ctx context.Context, sessionId string) (string, error)
	DeleteSession(ctx context.Context) (bool, error)
}

type Downloader struct{}
