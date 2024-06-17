package models

import (
	"github.com/gobuffalo/nulls"
	"time"
)

type Ytdl struct {
	Id        string     `json:"Id" db:"id"`
	CreatedAt time.Time  `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time  `json:"UpdatedAt" db:"updated_at"`
	DeletedAt nulls.Time `json:"DeletedAt" db:"deleted_at"`
	CreatedBy string     `json:"CreatedBy" db:"created_by"`
	Url       string     `json:"Url" db:"url"`
	StorePath string     `json:"StorePath" db:"store_path"`
	SessionId string     `json:"SessionId" db:"session_id"`
	Ttl       int        `json:"Ttl" db:"ttl"`
	Active    bool       `json:"Active" db:"active"`
	FileSize  int        `json:"FileSize" db:"file_size"`
}

func (Ytdl) TableName() string {
	return "ytdl"
}
