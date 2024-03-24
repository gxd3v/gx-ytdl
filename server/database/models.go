package database

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gx/youtubeDownloader/models"
)

type Database struct {
	DB *pop.Connection
}

type database interface {
	Transactional() *Database
	Commit()
	Rollback()
	List() ([]*any, error)
	Get(id string) (any, error)
	Insert(data ...any) (any, error)
	Update(data ...any) (any, error)
	Delete(data ...any) (any, error)
	GetByField(out any, conditions ...models.Condition) (any, error)
}
