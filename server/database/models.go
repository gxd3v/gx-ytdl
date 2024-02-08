package database

import (
	"gorm.io/gorm"
	"time"
)

type DB struct {
	Main *gorm.DB
}

type Ytdl struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy string
	Url       string
	StorePath string
	SessionId string
	Ttl       int
	Active    bool
	FileSize  int
}

type Log struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy string
	Message   string
	Arguments []interface{}
	Session   string
	Service   string
	Type      string
}

func LogTableName() string {
	return "logs"
}

func TableName() string {
	return "ytdl"
}

func ServiceTableName() string {
	return "services"
}

type Database interface {
	Connect(conn, table string) *DB
	Transactional(table string) *DB
	Insert(data interface{}, table string) error
	Delete(identifier string, data interface{}, table string) error
	GetAll(table string) ([]interface{}, error)
	Get(identifier string, table string) (*interface{}, error)
	Update(identifier string, data interface{}, table string) error
	Commit() error
	Rollback() error
}
