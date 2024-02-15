package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ database = (*Database)(nil)

func (db *Database) Connect(conn string) *Database {
	gormDB, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the main database " + err.Error())
	}

	db.Main = gormDB.Debug()

	return db
}

func (db *Database) Transactional() *Database {
	if db.Back == nil {
		db.Back = db.Main
	}
	db.Main = db.Back.Begin()
	return db
}

func (db *Database) Commit() error { return db.Main.Commit().Error }

func (db *Database) Rollback() error { return db.Main.Rollback().Error }

func (db *Database) Insert(model ...interface{}) error {
	for _, m := range model {
		err := db.Main.Model(m).Create(m).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Delete(model ...interface{}) error {
	for _, m := range model {
		err := db.Main.Model(m).Delete(m).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Get(id string) *Database {
	db.Main = db.Main.Where("id = ?", id)
	return db
}

func (db *Database) GetByField(field, value string) *Database {
	db.Main = db.Main.Where(fmt.Sprintf("%s = ?", field), value)
	return db
}

func (db *Database) List() []*interface{} {
	var data []*interface{}
	db.Main.Find(&data)

	return data
}

func (db *Database) Model(model interface{}) *Database {
	if db.Back == nil {
		db.Back = db.Main
	}
	db.Main = db.Back.Model(model)
	return db
}

func (db *Database) Scan(data interface{}) (interface{}, error) {
	err := db.Main.Scan(&data).Error
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("no data found")
	}
	return data, nil
}
