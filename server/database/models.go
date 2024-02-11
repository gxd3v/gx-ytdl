package database

import (
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Main *gorm.DB
}

type BaseModel struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy string
}

type Ytdl struct {
	BaseModel
	Url       string
	StorePath string
	SessionId string
	Ttl       int
	Active    bool
	FileSize  int
}

type BannedIP struct {
	BaseModel
	Ip string
}

func (BannedIP) TableName() string {
	return "banned_ips"
}

func (Ytdl) TableName() string {
	return "ytdl"
}

type database interface {
	Connect(conn string) *Database
	Transactional() *Database
	Commit() error
	Rollback() error
	Insert(model ...interface{}) error
	Delete(model ...interface{}) error
	Get(id string) *Database
	GetByField(field, value string) *Database
	List() []*interface{}
}
