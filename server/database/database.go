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

	db.DB = gormDB.Debug()
	return db
}

func (db *Database) Model(model any) *Database {
	db.DB = db.DB.Model(model)
	return db
}

func (db *Database) Transactional() *Database {
	return &Database{db.DB.Begin()}
}

func (db *Database) Commit() error { return db.DB.Commit().Error }

func (db *Database) Rollback() error { return db.DB.Rollback().Error }

func (db *Database) Insert(model ...interface{}) error {
	for _, m := range model {
		err := db.DB.Create(m).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Delete(model ...interface{}) error {
	for _, m := range model {
		err := db.DB.Delete(m).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Get(id string) *Database {
	db.DB = db.DB.Where("id = ?", id)
	return db
}

func (db *Database) GetByField(field, value string) *Database {
	db.DB = db.DB.Where(fmt.Sprintf("%s = ?", field), value)
	return db
}

func (db *Database) List() []*interface{} {
	var data []*interface{}
	db.DB.Find(&data)

	return data
}

func (db *Database) Scan(data interface{}) (interface{}, error) {
	err := db.DB.Scan(&data).Error
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("no data found")
	}
	return data, nil
}
