package models

import (
	"github.com/gobuffalo/nulls"
	"time"
)

type Session struct {
	Id        string     `json:"id" db:"id"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt nulls.Time `json:"deletedAt" db:"deleted_at"`
	CreatedBy string     `json:"createdBy" db:"created_by"`
	Session   string     `json:"session" db:"session"`
	LastLogin time.Time  `json:"lastLogin" db:"last_login"`
}

func (Session) TableName() string {
	return "session"
}
