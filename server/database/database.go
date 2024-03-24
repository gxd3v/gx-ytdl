package database

import (
	"fmt"
	"github.com/gobuffalo/pop/v6"
	"github.com/gx/youtubeDownloader/models"
)

var _ database = (*Database)(nil)

func Connect(env string) (*Database, error) {
	conn, err := pop.Connect(env)
	if err != nil {
		return nil, err
	}

	migrator, err := pop.NewFileMigrator("./migrations", conn)
	if err != nil {
		return nil, err
	}

	err = migrator.Up()
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: conn,
	}, nil
}

func (db *Database) Transactional() *Database {
	transaction, err := db.DB.Store.Transaction()
	if err != nil {
		return nil
	}

	db.DB.TX = transaction

	return db
}

func (db *Database) Commit() {
	_ = db.DB.TX.Commit()
}

func (db *Database) Rollback() {
	_ = db.DB.TX.Rollback()
}

func (db *Database) List() ([]*any, error) {
	var data []*any

	err := db.DB.All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *Database) Get(id string) (any, error) {
	var data any

	err := db.DB.Find(&data, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *Database) Insert(data ...any) (any, error) {
	err := db.DB.Create(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (db *Database) Update(data ...any) (any, error) {
	err := db.DB.Update(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (db *Database) Delete(data ...any) (any, error) {
	err := db.DB.Destroy(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (db *Database) GetByField(out any, conditions ...models.Condition) (any, error) {
	var query *pop.Query

	for _, condition := range conditions {
		query = query.Where(fmt.Sprintf("%v %v %v", condition.Field, condition.Operator, condition.Value))
	}

	err := query.All(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
