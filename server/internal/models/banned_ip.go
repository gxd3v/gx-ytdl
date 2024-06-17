package models

import (
	"github.com/gobuffalo/nulls"
	"time"
)

type BannedIP struct {
	Id        string     `json:"Id" db:"id"`
	CreatedAt time.Time  `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time  `json:"UpdatedAt" db:"updated_at"`
	DeletedAt nulls.Time `json:"DeletedAt" db:"deleted_at"`
	CreatedBy string     `json:"CreatedBy" db:"created_by"`
	Ip        string     `json:"Ip" db:"ip"`
}

func (BannedIP) TableName() string {
	return "banned_ip"
}
