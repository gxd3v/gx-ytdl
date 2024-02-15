package database

import (
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Main *gorm.DB
	Back *gorm.DB
}

//type BaseModel struct {
//	Id        string
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt time.Time
//	CreatedBy string
//}

type Ytdl struct {
	Id        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedBy string
	Url       string
	StorePath string
	SessionId string
	Ttl       int
	Active    bool
	FileSize  int
}

type BannedIP struct {
	Id        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedBy string
	Ip        string
}

type Session struct {
	Id        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedBy string
	Session   string
	LastLogin *time.Time
}

type Tabler interface {
	TableName() string
}

func (BannedIP) TableName() string {
	return "banned_ips"
}

func (Ytdl) TableName() string {
	return "store"
}

func (Session) TableName() string {
	return "sessions"
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

type controllers interface {
	NewSession() *Session
	NewYTDL(url string, storePath string, sessionId string, fileSize int) *Ytdl
}
