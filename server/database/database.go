package database

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ Database = (*DB)(nil)

func (db *DB) Connect(conn, table string) *DB {

	gormDB, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the main database " + err.Error())
	}

	db.Main = gormDB

	return db
}

func (db *DB) Transactional(table string) *DB {
	db.Main = db.Main.Table(table).Begin()
	return db
}

func (db *DB) Insert(data interface{}, table string) error {
	return db.Main.Table(table).Create(data).Error
}

func (db *DB) Delete(identifier string, data interface{}, table string) error {
	return db.Main.Table(table).Where("id = ?", identifier).Delete(&data).Error
}

func (db *DB) GetAll(table string) ([]interface{}, error) {
	var out []interface{}
	err := db.Main.Table(table).Find(&out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (db *DB) Get(identifier string, table string) (*interface{}, error) {
	data, ok := db.Main.Table(table).Get(identifier)
	if !ok {
		return nil, errors.New("failed to get record")
	}
	return &data, nil
}

func (db *DB) Update(identifier string, data interface{}, table string) error {
	return db.Main.Table(table).Where("id = ?", identifier).Updates(&data).Error
}

func (db *DB) Commit() error {
	return db.Main.Commit().Error
}

func (db *DB) Rollback() error {
	return db.Main.Rollback().Error
}
