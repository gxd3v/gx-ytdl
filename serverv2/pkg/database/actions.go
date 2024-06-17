package database

import (
	"fmt"
	"github.com/gobuffalo/pop/v6"
)

func New[T DBModel](db *pop.Connection) Generic[T] {
	return &generic[T]{
		db: db,
	}
}

func (g *generic[T]) DB() *pop.Connection {
	return g.db
}

func (g *generic[T]) Create(object ...*T) error {
	return g.db.Create(object)
}

func (g *generic[T]) Get(identifier string) (*T, error) {
	result := new(T)

	err := g.db.Where("id = ? and deleted_at = null", identifier).First(&result)

	return result, err
}

func (g *generic[T]) GetWithFilters(filters []*Filter) (*T, error) {
	var result = new(T)

	for _, filter := range filters {
		g.db.Where(fmt.Sprintf("%s %s ?", filter.Column, filter.Operator), filter.Value)
	}

	err := g.db.First(result)

	return result, err
}

func (g *generic[T]) List(filters []*Filter) ([]*T, int64, error) {
	result := make([]*T, 0)

	for _, filter := range filters {
		g.db.Where(fmt.Sprintf("%s %s ?", filter.Column, filter.Operator), filter.Value)
	}

	if err := g.db.All(&result); err != nil {
		return []*T{}, 0, err
	}

	return result, int64(len(result)), nil
}

func (g *generic[T]) Update(object ...*T) ([]*T, error) {
	if err := g.db.Update(object); err != nil {
		return nil, err
	}

	return object, nil
}

func (g *generic[T]) Delete(object ...*T) error {
	if err := g.db.Destroy(object); err != nil {
		return err
	}

	return nil
}
