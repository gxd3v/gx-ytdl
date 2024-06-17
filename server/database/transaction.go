package database

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gx/youtubeDownloader/internal/logger"
	"github.com/pkg/errors"
)

type TransactionManager interface {
	New() Transaction
}

type Transaction interface {
	DB() *pop.Connection
	Commit() error
	Rollback() error
}

type transactionManager struct {
	db *pop.Connection
}

type transaction struct {
	db *pop.Connection
}

func NewTransactionManager(db *pop.Connection) TransactionManager {
	return &transactionManager{
		db: db,
	}
}

func (m *transactionManager) Transaction(asd func(...any) error) error {
	return m.db.Transaction(func (tx *pop.Connection) error {
		return asd()
	})
}

func withTransaction(c buffalo.Context, txFunc func(tx *pop.Connection) error) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	if err := tx.Transaction(txFunc); err != nil {
		return errors.Wrap(err, "error in transaction")
	}

	return nil
}

func (t *transaction) DB() *pop.Connection {
	return t.db
}

func (t *transaction) Commit() error {
	return t.db.
}

func (t *transaction) Rollback() error {
	return t.db.Rollback()
}

//type TransactionManager interface {
//	New() Transaction
//}
//
//type Transaction interface {
//	DB() *gorm.DB
//	Commit() error
//	Rollback() error
//}
//
//type transactionManager struct {
//	db *gorm.DB
//}
//
//func NewTransactionManager(db *gorm.DB) TransactionManager {
//	return &transactionManager{
//		db: db,
//	}
//}
//
//func (m *transactionManager) New() Transaction {
//	return &transaction{
//		db: m.db.Begin(),
//	}
//}
//
//type transaction struct {
//	db *gorm.DB
//}
//
//func (t *transaction) DB() *gorm.DB {
//	return t.db
//}
//
//func (t *transaction) Commit() error {
//	return t.db.Commit().Error
//}
//
//func (t *transaction) Rollback() error {
//	return t.db.Rollback().Error
//}
