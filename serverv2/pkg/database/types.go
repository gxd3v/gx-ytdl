package database

import (
	"github.com/gobuffalo/pop/v6"
)

var _ Generic[DBModel] = (*generic[DBModel])(nil)

type DBModel interface {
	TableName() string
}

type generic[T DBModel] struct {
	db *pop.Connection
}

type Generic[T DBModel] interface {
	DB() *pop.Connection
	Create(object ...*T) error
	Get(identifier string) (*T, error)
	GetWithFilters(filters []*Filter) (*T, error)
	List(filters []*Filter) ([]*T, int64, error)
	Update(object ...*T) ([]*T, error)
	Delete(object ...*T) error
}

type Filter struct {
	Column   string `json:"column"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}
